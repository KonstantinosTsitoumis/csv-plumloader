package constants

import "ReadCsvCli/app/utils"

// Arguments Router

const HelpText = "███████████████████████████████████████████████████████████████████████████████████\n" +
	"█─▄▄▄─█─▄▄▄▄█▄─█─▄███▄─▄▄─█▄─▄███▄─██─▄█▄─▀█▀─▄█▄─▄███─▄▄─██▀▄─██▄─▄▄▀█▄─▄▄─█▄─▄▄▀█\n" +
	"█─███▀█▄▄▄▄─██▄▀▄█████─▄▄▄██─██▀██─██─███─█▄█─███─██▀█─██─██─▀─███─██─██─▄█▀██─▄─▄█\n" +
	"▀▄▄▄▄▄▀▄▄▄▄▄▀▀▀▄▀▀▀▀▀▄▄▄▀▀▀▄▄▄▄▄▀▀▄▄▄▄▀▀▄▄▄▀▄▄▄▀▄▄▄▄▄▀▄▄▄▄▀▄▄▀▄▄▀▄▄▄▄▀▀▄▄▄▄▄▀▄▄▀▄▄▀\n" +
	"-----------------------------------------------------------------------------------\n" +
	"Commands: \n" +
	"   upload       Upload Transactions CSV file to Database\n" +
	"-----------------------------------------------------------------------------------"
const UnknownCommand = "Unknown command, try reading the manual with 'help'"
const UploadCommandWronglyUsed = "To use upload command you need to provide a full csv file path\n " +
	"Example: upload user/example_directory/example_file.csv"

// Upload Function

const ReadingFile = "🤓\tReading file"
const ConvertingData = "🤓\tConverting data to database models..."

var FailedToReadFile = utils.Red("Failed to read file...")
var ConvertToDatabaseModelsErroredOrders = utils.Yellow("😟👉\tRows with TransactionID failed to be converted:\n⎿ ")
var UploadToDatabaseErroredOrders = utils.Yellow("💀\tRows with TransactionID failed to be Uploaded:\n⎿ ")
var UploadFinish = utils.Green("✅\tFinito!")
