package main

import (
	"fmt"
	"os"
)
func main(){
	text := "Hello world!"
	file, err := os.Create("C:\\GitHub\\MyStudy\\hello.txt")
	if err != nil{
		os.Exit(1)
		fmt.Println(err)
	}
	defer file.Close()

	file.WriteString(text)
}