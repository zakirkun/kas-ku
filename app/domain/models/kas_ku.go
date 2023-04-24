package models

import (
	"database/sql"
	"time"
)

type User struct {
	UserID        int          `db:"user_id" gorm:"primaryKey"`
	FullName      string       `db:"full_name"`
	Email         string       `db:"email"`
	PhoneNumber   string       `db:"phone"`
	Password      string       `db:"password"`
	LastIPLogin   string       `db:"last_ip_login"`
	LastLoginTime sql.NullTime `db:"last_login_time"`
	IsLocked      bool         `db:"is_locked" default:"false"`
	IsActive      bool         `db:"is_active" default:"false"`
	Pin           string       `db:"pin"`
	Token         string       `db:"token"`
	CreatedAt     time.Time
	UpdatedAt     time.Time
	DeletedAt     sql.NullTime `gorm:"index"`
}

type AccountInformation struct {
	IdDetail  int    `db:"id_detail" gorm:"primaryKey"`
	UserId    int    `db:"user_id"`
	FirstName string `db:"first_name"`
	LastName  string `db:"last_name"`
	Address   string `db:"address"`
	ZipCode   string `db:"zip_code"`
	Users     User   `gorm:"foreignKey:UserId"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt sql.NullTime `gorm:"index"`
}

type WalletCard struct {
	WalletID     int    `db:"wallet_id" gorm:"primaryKey"`
	UserId       int    `db:"user_id"`
	WalletNumber string `db:"wallet_number"`
	WalletTag    int    `db:"wallet_tag"`
	Balance      int64  `db:"balance"`
	IsLock       bool   `db:"is_lock"`
	Users        User   `gorm:"foreignKey:UserId"`
	CreatedAt    time.Time
	UpdatedAt    time.Time
	DeletedAt    sql.NullTime `gorm:"index"`
}

type WalletTag struct {
	TagId     int    `db:"tag_id" gorm:"primaryKey"`
	UserId    int    `db:"user_id"`
	WalletId  int    `db:"wallet_id"`
	TagName   string `db:"tag_name"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt sql.NullTime `gorm:"index"`
	Users     User         `gorm:"foreignKey:UserId"`
	Wallet    WalletCard   `gorm:"foreignKey:WalletId"`
}

type Security struct {
	Id         int    `db:"id" gorm:"primaryKey"`
	UserId     int    `db:"user_id"`
	Type       string `db:"type"`
	Content    string `db:"content"`
	ActionNeed string `db:"action_need"`
	Ip         string `db:"ip"`
	CreatedAt  time.Time
	UpdatedAt  time.Time
	DeletedAt  sql.NullTime `gorm:"index"`
	Users      User         `gorm:"foreignKey:UserId"`
}

type Notification struct {
	TagId     int    `db:"tag_id" gorm:"primaryKey"`
	UserId    int    `db:"user_id"`
	Title     string `db:"title"`
	Subject   string `db:"subject"`
	Message   string `db:"message"`
	IsOpen    bool   `db:"isOpen"`
	Sender    string `db:"sender"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt sql.NullTime `gorm:"index"`
	Users     User         `gorm:"foreignKey:UserId"`
}

type TransactionHistory struct {
	LogID     int    `db:"log_id" gorm:"primaryKey"`
	UserId    int    `db:"user_id"`
	LogType   string `db:"log_type"`
	Message   string `db:"message"`
	Amount    string `db:"amount"`
	Status    string `db:"status"`
	Users     User   `gorm:"foreignKey:UserId"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt sql.NullTime `gorm:"index"`
}
