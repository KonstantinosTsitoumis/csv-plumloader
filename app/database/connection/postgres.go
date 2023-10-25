package connection

import (
	"ReadCsvCli/app/database/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var modelsToMigrate = []interface{}{
	models.PlumTransaction{},
}

var makeGormToNotLog = logger.Default.LogMode(logger.Silent)

func NewDBConnection(dbUrl string) gorm.DB {
	db, connectionErr := gorm.Open(postgres.Open(dbUrl), &gorm.Config{
		Logger: makeGormToNotLog,
	})
	if connectionErr != nil {
		panic("Could not connect to database")
	}

	migrationErr := db.AutoMigrate(modelsToMigrate...)
	if migrationErr != nil {
		panic("Migration failed")
	}

	return *db
}
