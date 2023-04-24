package contracts

import (
	"github.com/labstack/echo"
	"github.com/zakirkun/kas-ku/app/domain/models"
	"github.com/zakirkun/kas-ku/app/domain/types"
)

type KasKuUsersRepository interface {
	CreateUsers(data models.User) error
	FindUsersByEmail(email string) bool
	CheckAccountLocked(email string) bool
	FindByToken(token string) *models.User
	UpdateUsers(UserID int, data models.User) error
}

type KasKuUsersServices interface {
	RegisterUsers(request types.UsersRegisterRequest) (error, *types.UsersRegisterResponse)
	ActivationUsers(request types.UsersActivationRequest) (error, *types.UsersActivationResponse)
}

type KasKuUsersDelivery interface {
	Register(e echo.Context) error
	Activation(e echo.Context) error
}
