package migration

import (
	"github.com/zakirkun/kas-ku/app/domain/models"
	"github.com/zakirkun/kas-ku/logger"
	"gorm.io/gorm"
)

func Migrate(db *gorm.DB) {

	err := db.AutoMigrate(&models.User{}, &models.AccountInformation{}, &models.WalletCard{}, &models.WalletTag{}, &models.Security{}, &models.Notification{}, &models.TransactionHistory{})

	if err != nil {
		logger.Logger.Info().Str("ERROR", err.Error()).Msg("Info migrations")
	}
}
