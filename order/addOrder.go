package order

import (
	"encoding/json"
	"fmt"
	"net/http"
)

const fileName string = "Orders.txt"

var statusServerError = 500

type DataBase interface {
	AddOrder(w http.ResponseWriter, r *http.Request)
	CancelOrder(w http.ResponseWriter, r *http.Request)
	ConfirmOrder(w http.ResponseWriter, r *http.Request)
	GetStatus(w http.ResponseWriter, r *http.Request)
	FindIdAndEditStatus(OrderDataBase []Order, intId int, statusOrder string) error
}

// сама структура базы данных
type Order struct {
	Product string
	Id      int
	Status  string
}

// реализация интерфейса базы данных
type InMemoryDataBase struct {
	data []Order
}

// конструктор базы данных
func NewInMemoryDataBase(Ord []Order) *InMemoryDataBase {
	return &InMemoryDataBase{
		data: Ord,
	}
}

// синглтон(надо как-то избегать)
var OrderDataBase DataBase = NewInMemoryDataBase([]Order{})

// главные интерфейсные функции, которые скрывают реализацию, но запускают процесс
func AddOrder(w http.ResponseWriter, r *http.Request) {
	OrderDataBase.AddOrder(w, r)
}

func CancelOrder(w http.ResponseWriter, r *http.Request) {
	OrderDataBase.CancelOrder(w, r)
}

func ConfirmOrder(w http.ResponseWriter, r *http.Request) {
	OrderDataBase.ConfirmOrder(w, r)
}

func GetStatus(w http.ResponseWriter, r *http.Request) {
	OrderDataBase.GetStatus(w, r)
}

func (db *InMemoryDataBase) AddOrder(w http.ResponseWriter, r *http.Request) {
	//первая часть - распаковка данных
	product := r.URL.Query().Get("order")

	bytesFile, err := getBytesFromFile(fileName)
	if err != nil {
		w.WriteHeader(statusServerError)
	}

	err = json.Unmarshal(bytesFile, &db.data)
	if err != nil {
		w.WriteHeader(statusServerError)
	}
	//вторая часть - работа с данными
	if len(db.data) > 0 {
		lastID := db.data[len(db.data)-1].Id
		db.data = append(db.data, Order{
			Status:  "ok",
			Id:      lastID + 1,
			Product: product,
		})
	} else {
		db.data = append(db.data, Order{
			Status:  "ok",
			Id:      1,
			Product: product,
		})
	}
	//третья часть - обратная запись данных в базу
	bytesOrder, err := json.Marshal(db.data)
	if err != nil {
		w.WriteHeader(statusServerError)
	}

	err = writeTextInFile(fileName, bytesOrder)
	if err != nil {
		w.WriteHeader(statusServerError)
	}

	fmt.Fprint(w, "Order added")
}
