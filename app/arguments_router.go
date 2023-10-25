package app

import (
	"ReadCsvCli/app/constants"
	"fmt"
	"strings"
)

type ArgumentsRouter struct {
	Upload func(string)
}

func NewArgumentsRouter(upload func(string)) ArgumentsRouter {
	return ArgumentsRouter{
		Upload: upload,
	}
}

func (a ArgumentsRouter) Run(arguments []string) {
	if len(arguments) == 1 {
		fmt.Println("Please provide a command or use 'help' to read manual.")
		return
	}

	command := strings.ToLower(arguments[1])

	restArguments := arguments[2:]

	switch command {
	case "upload":
		RunUpload(a.Upload, restArguments)
	case "help":
		fmt.Println(constants.HelpText)
	default:
		fmt.Println(constants.UnknownCommand)
	}
}

func RunUpload(uploadFunction func(string), arguments []string) {
	if len(arguments) != 1 || arguments[0] == "-h" {
		fmt.Println(constants.UploadCommandWronglyUsed)
		return
	}

	csvFilePath := arguments[0]

	uploadFunction(csvFilePath)
}
