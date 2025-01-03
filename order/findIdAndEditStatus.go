package order

import (
	"fmt"
)

func (db InMemoryDataBase) FindIdAndEditStatus(OrderDataBase []Order, intId int, statusOrder string) error {

	for i := 0; i < len(OrderDataBase); i++ {
		if OrderDataBase[i].Id == intId {
			OrderDataBase[i].Status = statusOrder
			return fmt.Errorf("")
		}
	}
	return fmt.Errorf("всё плохо")
}
