package order

import (
	"encoding/json"
	"fmt"
	"net/http"
)

const fileName string = "Orders.txt"

var statusServerError = 500

type Order struct {
	Product string
	Id      int
	Status  string
}

var OrderDataBase []Order

func AddOrder(w http.ResponseWriter, r *http.Request) {
	//первая часть - распаковка данных
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
