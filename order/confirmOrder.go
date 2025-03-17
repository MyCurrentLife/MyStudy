package order

import (
	"errors"
	"strconv"
)

func (db *FileDataBase) ConfirmOrder(id string) error {
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
	err = db.FindIdAndEditStatus(db.data, intId, "Confirm")
	if err.Error() == "всё плохо" {
		err = errors.New("id is missing")
		return err
	}

	//третья часть - обратная запись данных в базу
	err = db.writeDataBaseInFile()
	if err != nil {
		return errors.New(statusServerError)
	}
	return errors.New("product confirmed")
}
