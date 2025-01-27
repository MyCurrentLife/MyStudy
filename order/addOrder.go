package order

import (
	"encoding/json"
)

const fileName string = "Orders.txt"

var statusServerError = 500

func (db *InMemoryDataBase) AddOrder(product string) error {
	//первая часть - распаковка данных

	bytesFile, err := getBytesFromFile(fileName)
	if err != nil {
		return err
	}

	err = json.Unmarshal(bytesFile, &db.data)
	if err != nil {
		return err
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
		return err
	}

	err = writeTextInFile(fileName, bytesOrder)
	if err != nil {
		return err
	}
	return nil
}
