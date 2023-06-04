package gotest

import "fmt"

func Weather(location string) string {
	message := fmt.Sprintf("The weather looks nice today in %v!", location)
	return message
}
