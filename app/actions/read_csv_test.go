package actions

import (
	"encoding/csv"
	"fmt"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestReadCsv(t *testing.T) {
	// Arrange
	tearDownTest, filename := setupTest(t)
	defer tearDownTest(t)

	expected := []map[string]string{
		{
			"user_id":    "1",
			"first_name": "Ginny",
			"last_name":  "Benterman",
			"balance":    "5000",
		},
		{
			"user_id":    "2",
			"first_name": "Ricky",
			"last_name":  "McCold",
			"balance":    "10004",
		},
	}

	// Act
	result, err := ReadCsv(filename)

	// Assert
	assert.Equal(t, expected, result, "read csv result is not equal to expected")
	assert.Equal(t, nil, err, "err should have be nil")
}

// Test Helpers

func setupTest(tb testing.TB) (func(tb testing.TB), string) {
	MOCK_CSV_DATA := [][]string{
		{"user_id", "first_name", "last_name", "balance"},
		{"1", "Ginny", "Benterman", "5000"},
		{"2", "Ricky", "McCold", "10004"},
	}

	filename := createCSVFixture(tb, MOCK_CSV_DATA)

	return func(tb testing.TB) {
		err := os.Remove(filename)

		if err != nil {
			tb.Error("Test teardown failed")
		}
	}, filename
}

func createCSVFixture(tb testing.TB, content [][]string) string {
	filename := "test.csv"
	fmt.Println(filename)
	csvFile, err := os.Create(filename)

	if err != nil {
		tb.Error("Create csv fixture failed to create csv", err)
	}

	csvwriter := csv.NewWriter(csvFile)

	for _, row := range content {
		_ = csvwriter.Write(row[:])
	}
	csvwriter.Flush()

	return filename
}
