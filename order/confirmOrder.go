package order

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

func ConfirmOrder(w http.ResponseWriter, r *http.Request) {
	//первая часть - распаковка данных
	id := r.URL.Query().Get("id")

	intId, err := strconv.Atoi(id)
	if err != nil {
		w.WriteHeader(statusServerError)
	}

	b, err := getBytesFromFile(fileName)
	if err != nil {
		w.WriteHeader(statusServerError)
	}

	err = json.Unmarshal(b, &OrderDataBase)
	if err != nil {
		w.WriteHeader(statusServerError)
	}
	//вторая часть - работа с данными
	err = FindIdAndEditStatus(OrderDataBase, intId, "Confirm")
	if err.Error() == "Всё плохо!" {
		fmt.Fprintf(w, "id is missing")
	} else {
		fmt.Fprintf(w, "Product confirmed")
	}

	//третья часть - обратная запись данных в базу
	bytesorder, err := json.Marshal(OrderDataBase)
	if err != nil {
		w.WriteHeader(statusServerError)
	}

	err = writeTextInFile(fileName, bytesorder)
	if err != nil {
		w.WriteHeader(statusServerError)
	}
}
