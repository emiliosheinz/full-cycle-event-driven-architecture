package main

import (
	"context"
	"database/sql"
	"fmt"

	ckafka "github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/emiliosheinz/fc-ms-wallet-core/internal/database"
	event "github.com/emiliosheinz/fc-ms-wallet-core/internal/event"
	"github.com/emiliosheinz/fc-ms-wallet-core/internal/event/handler"
	createaccount "github.com/emiliosheinz/fc-ms-wallet-core/internal/usecase/create_account"
	createclient "github.com/emiliosheinz/fc-ms-wallet-core/internal/usecase/create_client"
	createtransaction "github.com/emiliosheinz/fc-ms-wallet-core/internal/usecase/create_transaction"
	"github.com/emiliosheinz/fc-ms-wallet-core/internal/web"
	"github.com/emiliosheinz/fc-ms-wallet-core/internal/web/webserver"
	"github.com/emiliosheinz/fc-ms-wallet-core/package/events"
	"github.com/emiliosheinz/fc-ms-wallet-core/package/kafka"
	"github.com/emiliosheinz/fc-ms-wallet-core/package/uow"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", "root", "root", "wallet-mysql", "3306", "wallet"))
	if err != nil {
		panic(err)
	}

	defer db.Close()

	_, err = db.Exec("CREATE TABLE IF NOT EXISTS clients (id varchar(255), name varchar(255), email varchar(255), created_at date, PRIMARY KEY (id))")
	if err != nil {
		panic(err)
	}
	_, err = db.Exec("CREATE TABLE IF NOT EXISTS accounts (id varchar(255), client_id varchar(255), balance integer, created_at date, PRIMARY KEY (id))")
	if err != nil {
		panic(err)
	}
	_, err = db.Exec("CREATE TABLE IF NOT EXISTS transactions (id varchar(255), account_id_from varchar(255), account_id_to varchar(255), amount integer, created_at date)")
	if err != nil {
		panic(err)
	}
	_, err = db.Exec("INSERT IGNORE INTO clients (id, name, email, created_at) VALUES ('1', 'Emilio Heinzmann', 'emiliosheinz@gmail.com', '2021-09-01')")
	if err != nil {
		panic(err)
	}
	_, err = db.Exec("INSERT IGNORE INTO accounts (id, client_id, balance, created_at) VALUES ('1', '1', 1000, '2021-09-01')")
	if err != nil {
		panic(err)
	}
	_, err = db.Exec("INSERT IGNORE INTO clients (id, name, email, created_at) VALUES ('2', 'John Doe', 'john@doe.com', '2021-09-01')")
	if err != nil {
		panic(err)
	}
	_, err = db.Exec("INSERT IGNORE INTO accounts (id, client_id, balance, created_at) VALUES ('2', '2', 1000, '2021-09-01')")
	if err != nil {
		panic(err)
	}

	configMap := ckafka.ConfigMap{
		"bootstrap.servers": "kafka:29092",
		"group.id":          "wallet",
	}
	kafkaProducer := kafka.NewKafkaProducer(&configMap)

	eventDispatcher := events.NewEventDispatcher()
	eventDispatcher.Register("TransactionCreated", handler.NewTransactionCreatedKafkaHandler(kafkaProducer))
	eventDispatcher.Register("BalanceUpdated", handler.NewBalanceUpdatedKafkaHandler(kafkaProducer))
	transactionCreatedEvent := event.NewTransactionCreated()
	balanceUpdatedEvent := event.NewBalanceUpdated()

	clientDb := database.NewClientDB(db)
	accountDb := database.NewAccountDB(db)

	ctx := context.Background()
	uow := uow.NewUow(ctx, db)

	uow.Register("AccountDB", func(tx *sql.Tx) interface{} {
		return database.NewAccountDB(db)
	})
	uow.Register("TransactionDB", func(tx *sql.Tx) interface{} {
		return database.NewTransactionDB(db)
	})

	createClientUseCase := createclient.NewCreateClientUseCase(clientDb)
	createAccountUseCase := createaccount.NewCreateAccountUseCase(accountDb, clientDb)
	createTransactionUseCase := createtransaction.NewCreateTransactionUseCase(uow, eventDispatcher, transactionCreatedEvent, balanceUpdatedEvent)

	webserver := webserver.NewWebServer(":3000")

	clientHandler := web.NewWebClientHandler(*createClientUseCase)
	accountHandler := web.NewWebAccountHandler(*createAccountUseCase)
	transactionHandler := web.NewWebTransactionHandler(*createTransactionUseCase)

	webserver.AddHandler("/clients", clientHandler.CreateClient)
	webserver.AddHandler("/accounts", accountHandler.CreateAccount)
	webserver.AddHandler("/transactions", transactionHandler.CreateTransaction)

	fmt.Println("🚀 Server running on port 3000")
	webserver.Start()
}
