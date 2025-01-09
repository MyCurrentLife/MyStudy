package order

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

func (db *InMemoryDataBase) CancelOrder(w http.ResponseWriter, r *http.Request) {
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

	_ = json.Unmarshal(b, &OrderDataBase)
	//вторая часть - работа с данными
	err = db.FindIdAndEditStatus(db.data, intId, "Cancel")
	if err.Error() == "Всё плохо!" {
		fmt.Fprintf(w, "id is missing")
	} else {
		fmt.Fprintf(w, "Product canceled")
	}
	//третья часть - обратная запись данных в базу
	bytesorder, err := json.Marshal(db.data)
	if err != nil {
		w.WriteHeader(statusServerError)
	}

	err = writeTextInFile(fileName, bytesorder)
	if err != nil {
		w.WriteHeader(statusServerError)
	}
}
