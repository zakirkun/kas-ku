package delivery

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/zakirkun/kas-ku/app/domain/contracts"
	"github.com/zakirkun/kas-ku/app/domain/types"
)

type deliveryUsersCtx struct {
	userServices contracts.KasKuUsersServices
}

func NewUsersDelivery(userServices contracts.KasKuUsersServices) contracts.KasKuUsersDelivery {
	return deliveryUsersCtx{userServices: userServices}
}

func (d deliveryUsersCtx) Activation(e echo.Context) error {
	var request types.UsersActivationRequest
	if err := e.Bind(&request); err != nil {
		return e.JSON(http.StatusBadRequest, types.ResponseApi{
			Code:    1,
			Message: "Bad Request",
		})
	}

	err, data := d.userServices.ActivationUsers(request)
	if err != nil {
		return e.JSON(http.StatusBadRequest, types.ResponseApi{
			Code:    1,
			Message: err.Error(),
		})
	}

	return e.JSON(http.StatusOK, types.ResponseApi{
		Code:    0,
		Message: "Activation Success.",
		Data:    data,
	})
}

func (d deliveryUsersCtx) Register(e echo.Context) error {

	var request types.UsersRegisterRequest
	if err := e.Bind(&request); err != nil {
		return e.JSON(http.StatusBadRequest, types.ResponseApi{
			Code:    1,
			Message: "Bad Request",
		})
	}

	err, data := d.userServices.RegisterUsers(request)
	if err != nil {
		return e.JSON(http.StatusBadRequest, types.ResponseApi{
			Code:    1,
			Message: err.Error(),
		})
	}

	return e.JSON(http.StatusCreated, types.ResponseApi{
		Code:    0,
		Message: "Register Success.",
		Data:    data,
	})
}
