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

const fileName string = "Orders.txt"
const statusOk int = 200
const statusServerError int = 500
const statusClientError int = 400

func main() {
	http.HandleFunc("/health", GetStatus)
	http.HandleFunc("/listOrder", GetDataBaseOrders)
	http.HandleFunc("/addOrder", AddOrder)
	http.HandleFunc("/confirmOrder", ConfirmOrder)
	http.HandleFunc("/cancelOrder", CancelOrder)

	hostName := ":5000"
	log.Fatal(http.ListenAndServe(hostName, nil))
}

func GetStatus(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "ok")
}

func GetDataBaseOrders(w http.ResponseWriter, r *http.Request) {
	bytesFile, err := getBytesFromFile(fileName)
	if err != nil {
		w.WriteHeader(statusServerError)
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(bytesFile)

}
func AddOrder(w http.ResponseWriter, r *http.Request) {
	//первая часть - распаковка данных
	var OrderDataBase []Order
	product := r.URL.Query().Get("order")

	bytesFile, err := getBytesFromFile(fileName)
	if err != nil {
		w.WriteHeader(statusServerError)
	}

	err = json.Unmarshal(bytesFile, &OrderDataBase)
	if err != nil {
		w.WriteHeader(statusServerError)
	}
	//вторая часть - работа с данными
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
	//третья часть - обратная запись данных в базу
	bytesOrder, err := json.Marshal(OrderDataBase)
	if err != nil {
		w.WriteHeader(statusServerError)
	}

	err = writeTextInFile(fileName, bytesOrder)
	if err != nil {
		w.WriteHeader(statusServerError)
	}

	fmt.Fprint(w, "Order added")
}
func ConfirmOrder(w http.ResponseWriter, r *http.Request) {
	//первая часть - распаковка данных
	var OrderDataBase []Order
	id := r.URL.Query().Get("id")

	intId, err := strconv.Atoi(id)
	if err != nil {
		w.WriteHeader(statusServerError)
	}

	b, err := getBytesFromFile(fileName)
	if err != nil {
		w.WriteHeader(statusServerError)
	}

	err = json.Unmarshal(b, &OrderDataBase)
	if err != nil {
		w.WriteHeader(statusServerError)
	}
	//вторая часть - работа с данными
	err = findIdAndEditStatus(OrderDataBase, intId, "Confirm")

	if err.Error() == "Всё плохо!" {
		fmt.Fprintf(w, "id is missing")
	} else {
		fmt.Fprintf(w, "Product confirmed")
	}

	//третья часть - обратная запись данных в базу
	bytesorder, err := json.Marshal(OrderDataBase)
	if err != nil {
		w.WriteHeader(statusServerError)
	}

	err = writeTextInFile(fileName, bytesorder)
	if err != nil {
		w.WriteHeader(statusServerError)
	}
}

func CancelOrder(w http.ResponseWriter, r *http.Request) {
	//первая часть - распаковка данных
	var OrderDataBase []Order
	id := r.URL.Query().Get("id")

	intId, err := strconv.Atoi(id)
	if err != nil {
		w.WriteHeader(statusServerError)
	}

	b, err := getBytesFromFile(fileName)
	if err != nil {
		w.WriteHeader(statusServerError)
	}

	err = json.Unmarshal(b, &OrderDataBase)
	//вторая часть - работа с данными
	err = findIdAndEditStatus(OrderDataBase, intId, "Cancel")
	if err.Error() == "Всё плохо!" {
		fmt.Fprintf(w, "id is missing")
	} else {
		fmt.Fprintf(w, "Product canceled")
	}
	//третья часть - обратная запись данных в базу
	bytesorder, err := json.Marshal(OrderDataBase)
	if err != nil {
		w.WriteHeader(statusServerError)
	}

	err = writeTextInFile(fileName, bytesorder)
	if err != nil {
		w.WriteHeader(statusServerError)
	}
}
func getBytesFromFile(name string) ([]byte, error) {
	b := make([]byte, 0, 0)
	file, err := os.OpenFile(name, os.O_RDWR, 0644)
	if err != nil {
		return b, err
	}

	defer file.Close()
	fileinfo, err := file.Stat()
	if err != nil {
		return b, err
	}

	filesize := fileinfo.Size()

	bytesFile := make([]byte, filesize)
	_, err = file.Read(bytesFile)
	return bytesFile, nil
}
func writeTextInFile(name string, b []byte) error {
	file, err := os.OpenFile(name, os.O_RDWR, 0644)
	if err != nil {
		return err
	}
	defer file.Close()
	file.Write(b)
	return nil
}
func findIdAndEditStatus(OrderDataBase []Order, intId int, statusOrder string) error {

	for i := 0; i < len(OrderDataBase); i++ {
		if OrderDataBase[i].Id == intId {
			OrderDataBase[i].Status = statusOrder
			return fmt.Errorf("Всё хорошо!")
		}
	}
	return fmt.Errorf("Всё плохо!")
}
