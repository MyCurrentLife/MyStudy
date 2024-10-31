package order

import (
	"reflect"
	"testing"
)

func TestFindIdAndEditStatus(t *testing.T) {
	//первый сценарий - хорошая концовка
	//входные данные

	var ord []Order = []Order{{Product: "banana", Id: 1, Status: "ok"}}

	var inputId = 1
	var inputStatus = "changed"

	var outputValue []Order = []Order{{Product: "banana", Id: 1, Status: "changed"}}

	//тест функции

	err := FindIdAndEditStatus(ord, inputId, inputStatus)
	if err.Error() == "Всё плохо!" {
		t.Errorf("error findIdAndEditStatus")
	}
	//проверка результатов

	if reflect.DeepEqual(ord, outputValue) == false {
		t.Errorf("value not a qual")
	}

	//второй сценарий - плохая концовка
	//входные данные

	inputId = 2
	inputStatus = "changed"

	//тест функции и проверка ошибки
	err = FindIdAndEditStatus(ord, inputId, inputStatus)
	if err.Error() != "всё плохо" {
		t.Errorf("error findIdAndEditStatus")
	}

}
