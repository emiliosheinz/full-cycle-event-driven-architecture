package main

import (
	"database/sql"
	"fmt"

	"github.com/emiliosheinz/fc-ms-wallet-core/internal/database"
	event "github.com/emiliosheinz/fc-ms-wallet-core/internal/events"
	createaccount "github.com/emiliosheinz/fc-ms-wallet-core/internal/usecase/create_account"
	createclient "github.com/emiliosheinz/fc-ms-wallet-core/internal/usecase/create_client"
	createtransaction "github.com/emiliosheinz/fc-ms-wallet-core/internal/usecase/create_transaction"
	"github.com/emiliosheinz/fc-ms-wallet-core/package/events"
)

func main() {
	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", "root", "root", "wallet-mysql", "3306", "wallet"))
	if err != nil {
		panic(err)
	}
	defer db.Close()
	eventDispatcher := events.NewEventDispatcher()
	transactionCreatedEvent := event.NewTransactionCreated()
	// eventDispatcher.Register()

	clientDb := database.NewClientDB(db)
	accountDb := database.NewAccountDB(db)
	transactionDb := database.NewTransactionDB(db)

	createClientUseCase := createclient.NewCreateClientUseCase(clientDb)
	createAccountUseCase := createaccount.NewCreateAccountUseCase(accountDb, clientDb)
	createTransactionUseCase := createtransaction.NewCreateTransactionUseCase(transactionDb, accountDb, eventDispatcher, transactionCreatedEvent)
}
