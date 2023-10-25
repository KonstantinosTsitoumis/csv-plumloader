package actions

import (
	"ReadCsvCli/app/database/models"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestConvertMapToDatabaseModel(t *testing.T) {
	// Arrange
	csvData := []map[string]string{
		{
			"Amount":                     "20.10",
			"CedingScheme":               "",
			"InvestorReference":          "e3f7c4c5-e575-4a76-9d24-45e1be2277dd",
			"FundName":                   "Cash",
			"Product":                    "ISA",
			"SEDOL":                      "Cash",
			"ISIN":                       "Cash",
			"TransactionDate":            "29/09/2023 00:00:00.0000000  +00:00",
			"TransactionType":            "Partial Withdrawal",
			"TransactionTypeID":          "21",
			"Units":                      "",
			"ClientTransactionReference": "4247E3AB-E2FB-4E68-AE1A-1CE0414E7E87",
			"iDate":                      "2023-09-29 05:00:07",
			"TransactionID":              "51975753",
		},
	}

	expectedIDate, _ := time.Parse(layoutIDate, "2023-09-29 05:00:07")
	expectedTransactionDate, _ := time.Parse(layoutTransactionDate, "29/09/2023 00:00:00.0000000  +00:00")

	expectedClientTransactionReference, _ := uuid.Parse("4247E3AB-E2FB-4E68-AE1A-1CE0414E7E87")
	expectedClientInvestorReference, _ := uuid.Parse("e3f7c4c5-e575-4a76-9d24-45e1be2277dd")

	expectedModel := models.PlumTransaction{
		Amount:                     2010,
		CedingScheme:               "",
		InvestorReference:          expectedClientInvestorReference,
		FundName:                   "Cash",
		Product:                    "ISA",
		Sedol:                      "Cash",
		ISIN:                       "Cash",
		TransactionDate:            expectedTransactionDate,
		TransactionType:            "Partial Withdrawal",
		TransactionTypeID:          21,
		Units:                      "",
		ClientTransactionReference: expectedClientTransactionReference,
		IDate:                      expectedIDate,
		TransactionID:              51975753,
	}

	// Act
	dbModels, errored := ConvertMapToDatabaseModel(csvData)

	// Assert
	model := dbModels[0]

	assert.Equal(t, expectedModel, model)
	assert.Equal(t, []string(nil), errored)
}
