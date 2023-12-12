package main

import (
	"fmt"
	"os"
)

func main() {
	text := "Hello world!"
	file, err := os.Create("hello.txt")
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	writeString, err := file.WriteString(text)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(writeString)
}
