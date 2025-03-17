package order

import (
	"errors"
)

const fileName string = "Orders.txt"

var statusServerError = "500"

func (db *FileDataBase) AddOrder(product string) error {
	//первая часть - распаковка данных
	var err error

	db.data, err = db.GetJSONFromFile()
	if err != nil {
		return errors.New(statusServerError)
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
	err = db.writeDataBaseInFile()
	if err != nil {
		return errors.New(statusServerError)
	}

	return errors.New("order added")
}
