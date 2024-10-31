package order

import (
	"fmt"
	"net/http"
)

func GetStatus(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "ok")
}
