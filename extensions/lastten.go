package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	fmt.Print("Main is up!\n")
}

func say_hello() string {
	fmt.Print("Hello world\n")
	return ("Hello world")
}

func sort(data string) string {
	var jsonData interface{}
	json.Unmarshal([]byte(data), &jsonData)
	fmt.Print(data, "\n")
	return ("Hello world")
}
