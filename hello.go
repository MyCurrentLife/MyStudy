package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type Order struct{
	 Product string
	 Status string
	 Id int
}
func main(){
	var testOrder Order = Order{"Стул","ok",32}
	bytes, err := json.Marshal(testOrder)
	if err != nil{
		fmt.Println("errorCreateJsonBytes")
	}
	file,err := os.Create("bytes.txt")
	if err != nil{
		fmt.Println("errorCreateFile")
	}
	defer file.Close()
	file.Write(bytes)
}
