package repositories

import (
	"ReadCsvCli/app/database/models"
	"errors"
	"gorm.io/gorm"
)

type PlumTransactionRepo struct {
	Db gorm.DB
}

func (p PlumTransactionRepo) InsertTransaction(transactionPointer *models.PlumTransaction) error {
	result := p.Db.Create(transactionPointer)

	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return errors.New("duplicate TransactionID")
	}

	return nil
}
