package order

import (
	"reflect"
	"testing"
)

type testCase struct {
	name        string
	dataBase    []Order
	inputId     int
	inputStatus string
	outputError bool
	errMsg      string
	outputValue []Order
}

func TestFindIdAndEditStatus(t *testing.T) {
	//первый сценарий - хорошая концовка
	//входные данные
	testCases := []testCase{
		{
			name:        "success",
			dataBase:    []Order{{Product: "banana", Id: 1, Status: "ok"}},
			inputId:     1,
			inputStatus: "changed",
			outputError: false,
			errMsg:      "",
			outputValue: []Order{{Product: "banana", Id: 1, Status: "changed"}},
		},
		{
			name:        "failure",
			dataBase:    []Order{{Product: "banana", Id: 1, Status: "ok"}},
			inputId:     2,
			inputStatus: "changed",
			outputError: true,
			errMsg:      "всё плохо",
			outputValue: []Order{{Product: "banana", Id: 1, Status: "changed"}},
		},
	}

	//тест функции
	for _, d := range testCases {
		err := FindIdAndEditStatus(d.dataBase, d.inputId, d.inputStatus)

		//проверка на ошибку

		if err.Error() != d.errMsg {
			t.Errorf("the error case %v should be like this %v, the message should be like this %v, but in fact it is like this %v", d.name, d.outputError, d.errMsg, err)
		}
		//проверка результатов

		if reflect.DeepEqual(d.dataBase, d.outputValue) == d.outputError {
			t.Errorf("%v case is not equal", d.name)
		}
	}
}
