package repositories

import (
	"ReadCsvCli/app/config"
	"ReadCsvCli/app/database/connection"
	"ReadCsvCli/app/database/models"
	"fmt"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
	"testing"
	"time"
)

func TestPlumTransactionRepo(t *testing.T) {
	db := getDBConnection()

	t.Cleanup(truncateDB(db))

	t.Run("test should insert PlumTransaction to database", testShouldInsertPlumTransaction(db))

}

func testShouldInsertPlumTransaction(db gorm.DB) func(t *testing.T) {
	return func(t *testing.T) {
		// Arrange
		plumTransactionRepo := PlumTransactionRepo{
			Db: db,
		}

		uuid_, _ := uuid.NewUUID()

		expectedTransaction := models.PlumTransaction{
			Amount:                     2010,
			InvestorReference:          uuid_,
			FundName:                   "test",
			Product:                    "test",
			Sedol:                      "test",
			ISIN:                       "test",
			TransactionDate:            time.Now(),
			TransactionType:            "test transaction",
			TransactionTypeID:          2,
			ClientTransactionReference: uuid_,
			IDate:                      time.Now(),
			TransactionID:              5555112312,
		}

		// Act
		_ = plumTransactionRepo.InsertTransaction(&expectedTransaction)

		// Assert
		var insertedTransaction models.PlumTransaction
		db.Where("transaction_id = ?", expectedTransaction.TransactionID).First(&insertedTransaction)

		assert.Equal(t, expectedTransaction.Amount, insertedTransaction.Amount)
		assert.Equal(t, expectedTransaction.InvestorReference, insertedTransaction.InvestorReference)
		assert.Equal(t, expectedTransaction.FundName, insertedTransaction.FundName)
		assert.Equal(t, expectedTransaction.Product, insertedTransaction.Product)
		assert.Equal(t, expectedTransaction.Sedol, insertedTransaction.Sedol)
		assert.Equal(t, expectedTransaction.ISIN, insertedTransaction.ISIN)
		assert.Equal(t, true, expectedTransaction.TransactionDate.Equal(insertedTransaction.TransactionDate))
		assert.Equal(t, expectedTransaction.TransactionType, insertedTransaction.TransactionType)
		assert.Equal(t, expectedTransaction.TransactionTypeID, insertedTransaction.TransactionTypeID)
		assert.Equal(t, expectedTransaction.ClientTransactionReference, insertedTransaction.ClientTransactionReference)
	}
}

// Helpers
func getDBConnection() gorm.DB {
	return connection.NewDBConnection(config.TEST_DB_URL)
}

func truncateDB(db gorm.DB) func() {
	return func() {
		var tables []string

		if err := db.Table("information_schema.tables").Where("table_schema = ?", "public").Pluck("table_name", &tables).Error; err != nil {
			panic(err)
		}

		for _, v := range tables {
			db.Exec(fmt.Sprintf("TRUNCATE TABLE %s;", v))
		}
	}
}
