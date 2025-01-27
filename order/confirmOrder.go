package order

import (
	"encoding/json"
	"errors"
	"strconv"
)

func (db *InMemoryDataBase) ConfirmOrder(id string) error {
	//первая часть - распаковка данных

	intId, err := strconv.Atoi(id)
	if err != nil {
		return err
	}

	b, err := getBytesFromFile(fileName)
	if err != nil {
		return err
	}

	err = json.Unmarshal(b, &OrderDataBase)
	if err != nil {
		return err
	}
	//вторая часть - работа с данными
	err = db.FindIdAndEditStatus(db.data, intId, "Confirm")
	if err.Error() == "Всё плохо!" {
		err = errors.New("id is missing")
		return err
	}

	//третья часть - обратная запись данных в базу
	bytesorder, err := json.Marshal(db.data)
	if err != nil {
		return err
	}

	err = writeTextInFile(fileName, bytesorder)
	if err != nil {
		return err
	}
	return nil
}
