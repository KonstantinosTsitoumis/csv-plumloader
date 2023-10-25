package actions

import (
	"ReadCsvCli/app/database/models"
	"github.com/google/uuid"
	"strconv"
	"time"
)

const layoutIDate = "2006-01-02 15:04:05"
const layoutTransactionDate = "02/01/2006 15:04:05.0000000  +00:00"

func ConvertMapToDatabaseModel(csvData []map[string]string) ([]models.PlumTransaction, []string) {
	var result []models.PlumTransaction
	var errored []string

	for _, row := range csvData {
		amount, amountErr := floatStringToMoneyValue(row["Amount"])
		investorReference, investorReferenceErr := uuid.Parse(row["InvestorReference"])
		transactionDate, transactionDateErr := time.Parse(layoutTransactionDate, row["TransactionDate"])
		iDate, iDateErr := time.Parse(layoutIDate, row["iDate"])
		transactionTypeID, transactionTypeIDErr := strconv.ParseInt(row["TransactionTypeID"], 0, 0)
		clientTransactionReference, clientTransactionReferenceErr := uuid.Parse(row["ClientTransactionReference"])
		transactionID, transactionIDErr := strconv.ParseInt(row["TransactionID"], 0, 0)

		if amountErr != nil ||
			investorReferenceErr != nil ||
			transactionDateErr != nil ||
			iDateErr != nil ||
			transactionTypeIDErr != nil ||
			clientTransactionReferenceErr != nil ||
			transactionIDErr != nil {
			errored = append(errored, row["TransactionID"])
			continue
		}

		result = append(
			result,
			models.PlumTransaction{
				Amount:                     amount,
				CedingScheme:               row["CedingScheme"],
				InvestorReference:          investorReference,
				FundName:                   row["FundName"],
				Product:                    row["Product"],
				Sedol:                      row["SEDOL"],
				ISIN:                       row["ISIN"],
				TransactionDate:            transactionDate,
				TransactionType:            row["TransactionType"],
				TransactionTypeID:          int(transactionTypeID),
				Units:                      row["Units"],
				ClientTransactionReference: clientTransactionReference,
				IDate:                      iDate,
				TransactionID:              int(transactionID),
			},
		)
	}

	return result, errored
}

func floatStringToMoneyValue(value string) (uint64, error) {
	const ToMoney = 100

	parsedAmount, err := strconv.ParseFloat(value, 32)

	if err != nil {
		return 0, err
	}

	intAmount := uint64(parsedAmount * ToMoney)

	return intAmount, nil
}
