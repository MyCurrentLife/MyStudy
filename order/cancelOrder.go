package order

import (
	"encoding/json"
	"errors"
	"strconv"
)

func (db *InMemoryDataBase) CancelOrder(id string) error {
	//первая часть - распаковка данных

	intId, err := strconv.Atoi(id)
	if err != nil {
		return err
	}

	b, err := getBytesFromFile(fileName)
	if err != nil {
		return err
	}

	_ = json.Unmarshal(b, &OrderDataBase)
	//вторая часть - работа с данными
	err = db.FindIdAndEditStatus(db.data, intId, "Cancel")
	if err.Error() == "Всё плохо!" {
		return errors.New("id is missing")
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
