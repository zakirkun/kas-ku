package services

import (
	"encoding/base64"
	"errors"
	"sync"
	"time"

	"github.com/zakirkun/kas-ku/config"

	"github.com/golang-jwt/jwt/v4"

	"github.com/zakirkun/kas-ku/app/domain/contracts"
	"github.com/zakirkun/kas-ku/app/domain/models"
	"github.com/zakirkun/kas-ku/app/domain/types"
	"github.com/zakirkun/kas-ku/helpers"
)

type usersServicesCtx struct {
	Repository contracts.KasKuUsersRepository
	Mutex      *sync.Mutex
}

func NewUsersServices(repo contracts.KasKuUsersRepository) contracts.KasKuUsersServices {
	return usersServicesCtx{Repository: repo, Mutex: &sync.Mutex{}}
}

func (s usersServicesCtx) ActivationPin(request types.PinActivationRequest) (error, bool) {
	// lock go routine
	s.Mutex.Lock()
	defer s.Mutex.Unlock()

	dataUsers := s.Repository.FindUsersByEmail(request.Email)
	if dataUsers == false {
		return errors.New("account not found"), false
	}

	updateUsers := models.User{
		Pin: request.PIN,
	}

	if err := s.Repository.UpdateUsers(request.UserID, updateUsers); err != nil {
		return errors.New("internal errors"), false
	}

	return nil, true
}

func (s usersServicesCtx) ActivationUsers(request types.UsersActivationRequest) (error, *types.UsersActivationResponse) {
	// lock go routine
	s.Mutex.Lock()
	defer s.Mutex.Unlock()

	dataUsers := s.Repository.FindByToken(request.Token)

	if dataUsers == nil {
		return errors.New("invalid activation"), nil
	}

	if dataUsers.IsActive {
		return errors.New("account is active"), nil
	}

	updateUsers := models.User{
		Token:    "NULL",
		IsActive: true,
	}

	if err := s.Repository.UpdateUsers(dataUsers.UserID, updateUsers); err != nil {
		return errors.New("internal errors"), nil
	}

	claims := &types.UsersClaims{
		dataUsers.UserID,
		dataUsers.Email,
		dataUsers.IsActive,
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 72)),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	t, _ := token.SignedString([]byte(config.GetString("JWT_SECRET")))

	return nil, &types.UsersActivationResponse{UserID: dataUsers.UserID, Type: "Bearer ", Token: t}
}

func (s usersServicesCtx) RegisterUsers(request types.UsersRegisterRequest) (error, *types.UsersRegisterResponse) {
	// lock go routine
	s.Mutex.Lock()
	defer s.Mutex.Unlock()

	// check email
	if isAvailable := s.Repository.FindUsersByEmail(request.Email); !isAvailable {
		return errors.New("email already taken"), nil
	}

	hashPassword, _ := helpers.HashPassword(request.Password)
	token := base64.StdEncoding.EncodeToString([]byte(helpers.GenerateToken(32)))

	dataUsers := models.User{
		FullName:    request.FullName,
		Email:       request.Email,
		PhoneNumber: request.PhoneNumber,
		Password:    hashPassword,
		IsLocked:    false,
		Token:       token,
		IsActive:    false,
	}

	if err := s.Repository.CreateUsers(dataUsers); err != nil {
		return errors.New("internal error"), nil
	}

	userReponse := types.UsersRegisterResponse{
		FullName: request.FullName,
		Email:    request.Email,
		Token:    token,
	}

	return nil, &userReponse
}
