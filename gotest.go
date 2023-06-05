package main

import (
	"fmt"
	"os"
)

func main() {
	for _, location := range os.Args[1:] {
		message := Weather(location)
		fmt.Println(message)
	}
}


func Weather(location string) string {
	message := fmt.Sprintf("The weather looks poor today in %v!", location)
	return message
}
