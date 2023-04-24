package database

import (
	"fmt"
	"os"

	"github.com/zakirkun/kas-ku/config"
	"github.com/zakirkun/kas-ku/logger"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func OpenDb() *gorm.DB {
	dsn := "%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local"

	db, err := gorm.Open(mysql.Open(fmt.Sprintf(dsn, config.GetString("DB_USER"), config.GetString("DB_PASS"), config.GetString("DB_HOST"), config.GetString("DB_PORT"), config.GetString("DB_NAME"))), &gorm.Config{})

	if err != nil {
		logger.Logger.Panic().Str("ERROR", err.Error()).Msg("Database Not Connected")
		os.Exit(1)

		return nil
	}

	return db
}
