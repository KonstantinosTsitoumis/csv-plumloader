package main

import (
	"ReadCsvCli/app"
	"os"
)

func main() {
	arguments := os.Args

	app.StartApp(arguments)
}
