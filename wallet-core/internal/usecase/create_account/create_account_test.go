package createaccount

import (
	"testing"

	"github.com/emiliosheinz/fc-ms-wallet-core/internal/entity"
	"github.com/emiliosheinz/fc-ms-wallet-core/internal/usecase/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestCreateAccountUseCase_Execute(t *testing.T) {
	client, _ := entity.NewClient("John Doe", "john@doe.com")
	clientGateway := &mocks.ClientGatewayMock{}
	clientGateway.On("Get", client.ID).Return(client, nil)

	accountGateway := &mocks.AccountGatewayMock{}
	accountGateway.On("Save", mock.Anything).Return(nil)

	uc := NewCreateAccountUseCase(accountGateway, clientGateway)
	inputDto := CreateAccountInputDTO{ClientID: client.ID}
	output, err := uc.Execute(inputDto)

	assert.Nil(t, err)
	assert.NotNil(t, output)
	assert.NotEmpty(t, output.ID)
	clientGateway.AssertExpectations(t)
	accountGateway.AssertExpectations(t)
	clientGateway.AssertNumberOfCalls(t, "Get", 1)
	accountGateway.AssertNumberOfCalls(t, "Save", 1)
}
