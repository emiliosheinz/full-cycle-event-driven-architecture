package gateway

import "github.com/emiliosheinz/fc-ms-wallet-core/internal/entity"

type AccountGateway interface {
	Save(account *entity.Account) error
	FindByID(id string) (*entity.Account, error)
}
