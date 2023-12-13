package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"os"
)

type OrderDataBase struct {
	SliceOrders []Order
}

type Order struct {
	Product  string
	Status   string
	Id       int
	SumFloat uint16
}

func main() {
	var Total OrderDataBase
	Total.SliceOrders = sliceRandomOrder(100)
	createFileWithText(Total)

	OrderBaseFromFile := CreateOrderDataBaseFromFile()
	MaxOrder := getMaxSumFloat(OrderBaseFromFile)
	fmt.Println(MaxOrder)
}

func CreateOrderDataBaseFromFile() OrderDataBase {
	content, err := os.ReadFile("Orders.txt")
	if err != nil {
		fmt.Println("errorReadFile")
	}

	var totalNew OrderDataBase

	err = json.Unmarshal(content, &totalNew.SliceOrders)
	if err != nil {
		fmt.Println("errorJsonUnmarshal")
	}
	return totalNew

}
func createFileWithText(total OrderDataBase) {
	bytesOrder, err := json.Marshal(total.SliceOrders)
	if err != nil {
		fmt.Println("errorJsonMarshal")
	}
	file, err := os.Create("Orders.txt")
	if err != nil {
		fmt.Println("errorCreateFile")
	}
	defer file.Close()
	file.Write(bytesOrder)
}

func sliceRandomOrder(n int) []Order {
	ord := []Order{}
	for i := 0; i < n; i++ {
		randomize := rand.Perm(1000)
		ord = append(ord, Order{})
		ord[i].Id = randomize[i]
		ord[i].SumFloat = uint16(randomize[i]) * 3
		ord[i].Status = "ok"
		ord[i].Product = "Стул"
	}
	return ord
}

func getMaxSumFloat(arr OrderDataBase) Order {
	max := arr.SliceOrders[0]
	for _, element := range arr.SliceOrders[:10] {
		if element.SumFloat > max.SumFloat {
			max = element
		}
	}
	return max
}
