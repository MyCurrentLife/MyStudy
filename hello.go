package main

import (
	"fmt"
	"os"
)
func main(){
	text := "Hello world!"
	file, err := os.Create("hello.txt")
	if err != nil{
		fmt.Println("error")
	}
	defer file.Close()

	file.WriteString(text)
}