package app

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestShouldCallUploadFunction(tb *testing.T) {
	// Arrange
	uploadSpy := NewArgumentsRouterSpy()

	argumentsRouter := NewArgumentsRouter(
		uploadSpy.spyFunc,
	)

	arguments := []string{"...", "upload", "test/test/test.csv"}

	expected := "test/test/test.csv"

	// Act
	argumentsRouter.Run(arguments)

	// Assert
	assert.Equal(tb, 1, uploadSpy.getCallCount())
	assert.Equal(tb, expected, uploadSpy.getArgsSlice()[0])
}

// Test Helpers

type ArgumentsRouterSpy struct {
	getCallCount func() int
	getArgsSlice func() []string
	spyFunc      func(string)
}

func NewArgumentsRouterSpy() ArgumentsRouterSpy {
	callCounter := 0
	callArgs := make([]string, 0)

	callCounterPointer := &callCounter
	callArgsPointer := &callArgs

	return ArgumentsRouterSpy{
		func() int {
			return *callCounterPointer
		},
		func() []string {
			return *callArgsPointer
		},
		func(filepath string) {
			*callCounterPointer++
			*callArgsPointer = append(*callArgsPointer, filepath)
		},
	}
}
