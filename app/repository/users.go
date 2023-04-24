package repository

import (
	"github.com/zakirkun/kas-ku/app/domain/contracts"
	"github.com/zakirkun/kas-ku/app/domain/models"
	"gorm.io/gorm"
)

type usersRepositoryCtx struct {
	DB *gorm.DB
}

func NewUsersRepository(db *gorm.DB) contracts.KasKuUsersRepository {
	return usersRepositoryCtx{DB: db}
}

func (r usersRepositoryCtx) CreateUsers(data models.User) error {

	err := r.DB.Create(&data).Error

	return err
}

func (r usersRepositoryCtx) FindUsersByEmail(email string) bool {
	var users models.User
	r.DB.Where("email = ?", email).Find(&users)

	return users.Email == ""
}

func (r usersRepositoryCtx) CheckAccountLocked(email string) bool {
	var users models.User
	r.DB.Where(models.User{
		IsLocked: true,
	}).Find(&users)

	return users.Email == ""
}

func (r usersRepositoryCtx) FindByToken(token string) *models.User {

	var users models.User
	r.DB.Where("token = ?", token).Find(&users)

	if users.UserID == 0 {
		return nil
	}

	return &users
}

func (r usersRepositoryCtx) UpdateUsers(UserID int, data models.User) error {

	return r.DB.Model(&models.User{}).Where("user_id = ?", UserID).Updates(data).Error
}
