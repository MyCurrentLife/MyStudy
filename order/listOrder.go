package order

import (
	"net/http"
)

func GetDataBaseOrders(w http.ResponseWriter, r *http.Request) {
	bytesFile, err := getBytesFromFile(fileName)
	if err != nil {
		w.WriteHeader(statusServerError)
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(bytesFile)
}
