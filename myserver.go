package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
)

type Order struct {
	Product string
	Id      int
	Status  string
}

var OrderDataBase []Order

func main() {
	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		GetStatus(w, *r)
	})
	http.HandleFunc("/listOrder", func(w http.ResponseWriter, r *http.Request) {
		GetDataBaseOrders(w, *r)
	})
	http.HandleFunc("/addOrder", func(w http.ResponseWriter, r *http.Request) {
		AddOrder(w, *r)
	})
	http.HandleFunc("/confirmOrder", func(w http.ResponseWriter, r *http.Request) {
		ConfirmOrder(w, *r)
	})
	http.HandleFunc("/cancelOrder", func(w http.ResponseWriter, r *http.Request) {
		CancelOrder(w, *r)
	})

	port := ":5000"
	log.Fatal(http.ListenAndServe(port, nil))
}

func GetStatus(w http.ResponseWriter, r http.Request) {
	fmt.Fprintf(w, "ok")
}

func GetDataBaseOrders(w http.ResponseWriter, r http.Request) {
	bytesFile := getBytesFromFile("Orders.txt",w)

	w.Header().Set("Content-Type", "application/json")
	w.Write(bytesFile)

}
func AddOrder(w http.ResponseWriter, r http.Request) {
	product := r.URL.Query().Get("order")

	bytesFile := getBytesFromFile("Orders.txt",w)
	
	err := json.Unmarshal(bytesFile, &OrderDataBase)

	if len(OrderDataBase) > 0 {
		lastID := OrderDataBase[len(OrderDataBase)-1].Id
		OrderDataBase = append(OrderDataBase, Order{
			Status:  "ok",
			Id:      lastID + 1,
			Product: product,
		})
	} else {
		OrderDataBase = append(OrderDataBase, Order{
			Status:  "ok",
			Id:      1,
			Product: product,
		})
	}

	bytesOrder, err := json.Marshal(OrderDataBase)
	if err != nil {
		fmt.Fprintf(w, "500")
	}

	writeTextInFile("Orders.txt", bytesOrder,w)

	fmt.Fprint(w, "Order added")
}
func ConfirmOrder(w http.ResponseWriter, r http.Request) {
	id := r.URL.Query().Get("id")
	
	intId, err := strconv.Atoi(id)
	if err != nil {
		fmt.Fprintf(w,"500")
	}

	b := getBytesFromFile("Orders.txt",w)

	err = json.Unmarshal(b, &OrderDataBase)
	
	for i := 0; i < len(OrderDataBase); i++{
		if OrderDataBase[i].Id == intId{
			OrderDataBase[i].Status = "Confirm"
			fmt.Fprintf(w, "Order confirmed")
		}
	}
	
	bytesorder, err := json.Marshal(OrderDataBase)
	if err != nil {
		fmt.Fprintf(w,"500")
	}
	
	writeTextInFile("Orders.txt",bytesorder,w)
}
func CancelOrder(w http.ResponseWriter, r http.Request) {
	id := r.URL.Query().Get("id")
	
	intId, err := strconv.Atoi(id)
	if err != nil {
		fmt.Fprintf(w,"500")
	}

	b := getBytesFromFile("Orders.txt", w)

	err = json.Unmarshal(b, &OrderDataBase)
	
	for i := 0; i < len(OrderDataBase)-1; i++{
		if OrderDataBase[i].Id == intId{
			OrderDataBase[i].Status = "Cancel"
			fmt.Fprintf(w, "Order canceled")
		}
	}
	
	bytesorder, err := json.Marshal(OrderDataBase)
	if err != nil {
		fmt.Fprintf(w,"500")
	}
	
	writeTextInFile("Orders.txt", bytesorder,w)
}
func getBytesFromFile(name string, w http.ResponseWriter)[]byte{
	file, err := os.OpenFile(name, os.O_RDWR, 0644)
	if err != nil {
		fmt.Fprintf(w,"500")
	}

	defer file.Close()
	fileinfo, err := file.Stat()
	if err != nil {
		fmt.Fprintf(w,"500")
	}

	filesize := fileinfo.Size()

	bytesFile := make([]byte, filesize)
	_, err = file.Read(bytesFile)
	return bytesFile
}
func writeTextInFile(name string, b []byte,w http.ResponseWriter){
	file,err := os.OpenFile(name, os.O_RDWR, 0644)
	if err != nil{
		fmt.Fprintf(w,"500")
	}


	defer file.Close()
	file.Write(b)
}
