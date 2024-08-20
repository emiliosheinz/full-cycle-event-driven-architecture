package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateNewTransaction(t *testing.T) {
	client1, _ := NewClient("John Doe", "john@doe.con")
	account1 := NewAccount(client1)
	client2, _ := NewClient("Jane Doe", "jane@doe.com")
	account2 := NewAccount(client2)

	account1.Credit(1000)
	account2.Credit(1000)

	transaction, error := NewTransaction(account1, account2, 100)
	assert.Nil(t, error)
	assert.NotNil(t, transaction)
	assert.Equal(t, 900.0, account1.Balance)
	assert.Equal(t, 1100.0, account2.Balance)
}

func TestCreateNewTransactionWithInsusfficientBalance(t *testing.T) {
	client1, _ := NewClient("John Doe", "john@doe.con")
	account1 := NewAccount(client1)
	client2, _ := NewClient("Jane Doe", "jane@doe.com")
	account2 := NewAccount(client2)

	account1.Credit(1000)
	account2.Credit(1000)

	transaction, error := NewTransaction(account1, account2, 1001)
	assert.NotNil(t, error)
	assert.Error(t, error, "insufficient balance")
	assert.Nil(t, transaction)
}

func TestCreateNewTransactionWithInvalidAmount(t *testing.T) {
	client1, _ := NewClient("John Doe", "john@doe.con")
	account1 := NewAccount(client1)
	client2, _ := NewClient("Jane Doe", "jane@doe.com")
	account2 := NewAccount(client2)

	account1.Credit(1000)
	account2.Credit(1000)

	transaction, error := NewTransaction(account1, account2, 0)
	assert.NotNil(t, error)
	assert.Error(t, error, "amount must be greater than zero")
	assert.Nil(t, transaction)
}
