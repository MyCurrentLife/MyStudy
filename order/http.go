package order

import (
	"fmt"
	"net/http"
)

type DataBase interface {
	AddOrder(product string) error
	CancelOrder(id string) error
	ConfirmOrder(id string) error
	GetStatus() (string, error)
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
	product := r.URL.Query().Get("order")

	err := OrderDataBase.AddOrder(product)
	if err != nil {
		w.WriteHeader(statusServerError)
	} else {
		fmt.Fprint(w, "Order added")
	}
}

func CancelOrder(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")

	err := OrderDataBase.CancelOrder(id)

	if err != nil && err.Error() != "id is missing" {
		w.WriteHeader(statusServerError)
	}
	if err.Error() == "id is missing" {
		fmt.Fprint(w, err.Error())
	}
	if err.Error() == "Product canceled" {
		fmt.Fprint(w, err.Error())
	}

}

func ConfirmOrder(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	err := OrderDataBase.ConfirmOrder(id)
	if err != nil && err.Error() != "id is missing" {
		w.WriteHeader(statusServerError)
	}
	if err.Error() == "id is missing" {
		fmt.Fprint(w, err.Error())
	}
	if err.Error() == "Product confirmed" {
		fmt.Fprint(w, err.Error())
	}

}

func GetStatus(w http.ResponseWriter, r *http.Request) {
	status, err := OrderDataBase.GetStatus()
	if err != nil {
		w.WriteHeader(statusServerError)
	}
	fmt.Fprint(w, status)
}
