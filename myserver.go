package main

import (
	"encoding/json"
	"fmt"
	"io"
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
	file, err := os.OpenFile("Orders.txt", os.O_RDWR, 0644)
	if err != nil {
		fmt.Fprintf(w, "500")
	}

	fileinfo, err := file.Stat()
	if err != nil {
		fmt.Fprintf(w, "500")
	}
	filesize := fileinfo.Size()

	bytesFile := make([]byte, filesize)

	_, err = file.Read(bytesFile)

	defer file.Close()
	w.Header().Set("Content-Type", "application/json")
	w.Write(bytesFile)

}
func AddOrder(w http.ResponseWriter, r http.Request) {
	product := r.URL.Query().Get("order")

	file, err := os.OpenFile("Orders.txt", os.O_RDWR, 0644)
	if err != nil {
		fmt.Fprintf(w, "500")
	}

	defer file.Close()
	fileinfo, err := file.Stat()
	if err != nil {
	}

	filesize := fileinfo.Size()

	bytesFile := make([]byte, filesize)
	_, err = file.Read(bytesFile)
	err = json.Unmarshal(bytesFile, &OrderDataBase)

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

	writeText, err := os.OpenFile("Orders.txt", os.O_RDWR, 0644)
	if err != nil {
		fmt.Fprintf(w, "500")
	}

	bytesOrder, err := json.Marshal(OrderDataBase)
	if err != nil {
		fmt.Fprintf(w, "500")
	}

	writeText.Write(bytesOrder)

	defer writeText.Close()

	fmt.Fprint(w, "Order added")
}
func ConfirmOrder(w http.ResponseWriter, r http.Request) {
	id := r.URL.Query().Get("id")
	intId, err := strconv.Atoi(id)
	if err != nil {
	}

	file, err := os.OpenFile("Orders.txt", os.O_RDWR, 0644)
	if err != nil {
	}
	b, err := io.ReadAll(file)
	if err != nil {
	}

	err = json.Unmarshal(b, &OrderDataBase)
	OrderDataBase[intId-1].Status = "Confirmed"
	bytesorder, err := json.Marshal(OrderDataBase)
	if err != nil {
	}
	
	writetext,err := os.OpenFile("Orders.txt", os.O_RDWR, 0644)
	if err != nil{}
	
	
	writetext.Write(bytesorder)


	defer file.Close()
	defer writetext.Close()


	fmt.Fprintf(w, "zaebalsa")
}
func CancelOrder(w http.ResponseWriter, r http.Request) {
	id := r.URL.Query().Get("id")
	intId, err := strconv.Atoi(id)
	if err != nil {
	}

	file, err := os.OpenFile("Orders.txt", os.O_RDWR, 0644)
	if err != nil {
	}
	b, err := io.ReadAll(file)
	if err != nil {
	}

	err = json.Unmarshal(b, &OrderDataBase)
	OrderDataBase[intId-1].Status = "Canceled"
	bytesorder, err := json.Marshal(OrderDataBase)
	if err != nil {
	}
	
	writetext,err := os.OpenFile("Orders.txt", os.O_RDWR, 0644)
	if err != nil{}
	
	
	writetext.Write(bytesorder)


	defer file.Close()
	defer writetext.Close()


	fmt.Fprintf(w, "zaebalsa")
}
