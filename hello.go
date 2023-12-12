package main

import "os"

func main(){
	text := "Hello world!"
	file, err := os.Create("hello.txt")

	if err != nil{
		os.Exit(1)
	}
	file.WriteString(text)
}