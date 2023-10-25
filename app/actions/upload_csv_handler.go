package actions

import (
	"ReadCsvCli/app/constants"
	"ReadCsvCli/app/database/models"
	"ReadCsvCli/app/database/repositories"
	"fmt"
	"github.com/schollz/progressbar/v3"
)

type readCsvFunction func(filename string) ([]map[string]string, error)
type convertMapToDatabaseModelFunction func([]map[string]string) ([]models.PlumTransaction, []string)

type UploadCsvHandlerDependencies struct {
	Repo                    repositories.PlumTransactionRepo
	ConvertToDatabaseModels convertMapToDatabaseModelFunction
	ReadCsv                 readCsvFunction
}

func UploadCsvHandler(
	deps UploadCsvHandlerDependencies,
) func(string) {
	return func(csvFilePath string) {
		fmt.Println("[Step 1/3]", constants.ReadingFile, csvFilePath)
		csvData, err := deps.ReadCsv(csvFilePath)

		if err != nil {
			fmt.Println(constants.FailedToReadFile, err)
			return
		}

		fmt.Println("[Step 2/3]", constants.ConvertingData)
		convertedData, erroredData := deps.ConvertToDatabaseModels(csvData)

		if len(erroredData) != 0 {
			fmt.Println(constants.ConvertToDatabaseModelsErroredOrders, erroredData)
		}

		bar := *progressbar.Default(
			int64(len(convertedData)),
			"[Step 3/3] 🚀      Uploading CSV rows to Database...",
		)
		defer bar.Close()

		erroredInserts := uploadToDatabase(deps.Repo, convertedData, &bar)

		if len(erroredInserts) != 0 {
			fmt.Println(constants.UploadToDatabaseErroredOrders, erroredInserts)
		}

		fmt.Println(constants.UploadFinish)
	}
}

func uploadToDatabase(repo repositories.PlumTransactionRepo, modelsSlice []models.PlumTransaction, bar *progressbar.ProgressBar) []string {
	var errored []string

	for _, v := range modelsSlice {
		bar.Add(1)

		err := repo.InsertTransaction(&v)

		if err != nil {
			errored = append(errored, string(string(rune(v.TransactionID))+string(err.Error())))
		}

	}

	return errored
}
