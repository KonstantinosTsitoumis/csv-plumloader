package actions

import (
	"encoding/csv"
	"errors"
	"os"
	"strings"
)

func ReadCsv(filename string) ([]map[string]string, error) {
	csvfile, err := os.Open(filename)
	defer csvfile.Close()

	if err != nil {
		return nil, errors.New("could not open file")
	}

	reader := csv.NewReader(csvfile)
	reader.Comma = ','

	rawCSVData, err := reader.ReadAll()

	if err != nil {
		return nil, errors.New("could not read csv")
	}

	return csvToMap(&rawCSVData), nil
}

func csvToMap(rawCSVDataPointer *[][]string) []map[string]string {
	var headers []string
	var csvMap []map[string]string

	rawCSVData := *rawCSVDataPointer

	headersRow := rawCSVData[0]

	for _, header := range headersRow {
		headers = append(headers, header)
	}

	for _, record := range rawCSVData[1:] {
		line := map[string]string{}

		for i, item := range record {
			line[headers[i]] = strings.Trim(item, "\u0000")
		}

		csvMap = append(csvMap, line)
	}

	return csvMap
}
