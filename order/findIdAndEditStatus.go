package order

import (
	"errors"
)

func (db *FileDataBase) FindIdAndEditStatus(OrderDataBase []Order, intId int, statusOrder string) error {

	if intId > len(OrderDataBase) {
		return errors.New("всё плохо")
	}

	for i := 0; i < len(OrderDataBase); i++ {
		if OrderDataBase[i].Id == intId {
			OrderDataBase[i].Status = statusOrder
		}
	}
	return errors.New("")
}
