package order

import (
	"fmt"
	"net/http"
)

func GetDataBaseOrders(w http.ResponseWriter, r *http.Request) {
	bytesFile, err := getBytesFromFile(fileName)
	if err != nil {
		fmt.Fprint(w, statusServerError)
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(bytesFile)
}
