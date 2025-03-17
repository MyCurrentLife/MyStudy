package order

import (
	"errors"
	"strconv"
)

func (db *FileDataBase) CancelOrder(id string) error {
	//первая часть - распаковка данных

	intId, err := strconv.Atoi(id)
	if err != nil {
		return err
	}

	db.data, err = db.GetJSONFromFile()
	if err != nil {
		return errors.New(statusServerError)
	}
	//вторая часть - работа с данными
	err = db.FindIdAndEditStatus(db.data, intId, "Cancel")
	if err.Error() == "всё плохо" {
		return errors.New("id is missing")
	}
	//третья часть - обратная запись данных в базу
	err = db.writeDataBaseInFile()
	if err != nil {
		return errors.New(statusServerError)
	}

	return errors.New("product Canceled")
}
