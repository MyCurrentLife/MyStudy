package order

import (
	"fmt"
	"net/http"
)

func (db *InMemoryDataBase) GetStatus(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "ok")
}
