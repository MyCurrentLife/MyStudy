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
type FileDataBase struct {
	data []Order
}

// конструктор базы данных
func NewFileDataBase(Ord []Order) *FileDataBase {
	return &FileDataBase{
		data: Ord,
	}
}

// синглтон(надо как-то избегать)
var OrderDataBase DataBase = NewFileDataBase([]Order{})

// главные интерфейсные функции, которые скрывают реализацию, но запускают процесс
func AddOrder(w http.ResponseWriter, r *http.Request) {
	product := r.URL.Query().Get("order")

	err := OrderDataBase.AddOrder(product)

	if err.Error() == "order added" {
		fmt.Fprint(w, err.Error())
	} else {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

}

func CancelOrder(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")

	err := OrderDataBase.CancelOrder(id)

	if err.Error() == "order Canceled" {
		fmt.Fprint(w, err.Error())
	} else {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func ConfirmOrder(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")

	err := OrderDataBase.ConfirmOrder(id)

	if err.Error() == "product confirmed" {
		fmt.Fprint(w, err.Error())
	} else {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func GetStatus(w http.ResponseWriter, r *http.Request) {
	status, err := OrderDataBase.GetStatus()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	fmt.Fprint(w, status)
}
