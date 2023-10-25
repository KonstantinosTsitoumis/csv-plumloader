package app

import (
	"ReadCsvCli/app/actions"
	"ReadCsvCli/app/config"
	"ReadCsvCli/app/database/connection"
	"ReadCsvCli/app/database/repositories"
)

func StartApp(cliArguments []string) {
	argumentsRouter := initialize(config.DB_URL)

	argumentsRouter.Run(cliArguments)
}

func initialize(dbUrl string) ArgumentsRouter {
	db := connection.NewDBConnection(dbUrl)

	repo := repositories.PlumTransactionRepo{Db: db}

	uploadCsvHandlerDependencies := actions.UploadCsvHandlerDependencies{
		Repo:                    repo,
		ConvertToDatabaseModels: actions.ConvertMapToDatabaseModel,
		ReadCsv:                 actions.ReadCsv,
	}

	uploadFunction := actions.UploadCsvHandler(uploadCsvHandlerDependencies)

	return NewArgumentsRouter(uploadFunction)
}
