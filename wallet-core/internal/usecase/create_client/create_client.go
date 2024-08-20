package createclient

import (
	"github.com/emiliosheinz/fc-ms-wallet-core/internal/entity"
	"github.com/emiliosheinz/fc-ms-wallet-core/internal/gateway"
)

type CreateClientDTO struct {
	Name  string
	Email string
}

type CreateClientOutputDTO struct {
	ID        string
	Name      string
	Email     string
	CreatedAt string
	UpdatedAt string
}

type CreateClientUseCase struct {
	ClientGatewat gateway.ClientGateway
}

func NewCreateClientUseCase(clientGateway gateway.ClientGateway) *CreateClientUseCase {
	return &CreateClientUseCase{ClientGatewat: clientGateway}
}

func (uc *CreateClientUseCase) Execute(dto *CreateClientDTO) (*CreateClientOutputDTO, error) {
	client, err := entity.NewClient(dto.Name, dto.Email)
	if err != nil {
		return nil, err
	}
	err = uc.ClientGatewat.Save(client)
	if err != nil {
		return nil, err
	}
	return &CreateClientOutputDTO{
		ID:        client.ID,
		Name:      client.Name,
		Email:     client.Email,
		CreatedAt: client.CreatedAt.String(),
		UpdatedAt: client.UpdatedAt.String(),
	}, nil
}
