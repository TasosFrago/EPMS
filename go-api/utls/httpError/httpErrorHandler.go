package httpError

import (
	"fmt"
	"net/http"
)

func InternalServerError(w http.ResponseWriter, errMsg string) {
	fmt.Print(errMsg)
	w.WriteHeader(http.StatusInternalServerError)
	w.Write([]byte("500: Internal Server Error"))
}

func NotFoundError(w http.ResponseWriter, errMsg string) {
	fmt.Print(errMsg)
	w.WriteHeader(http.StatusNotFound)
	w.Write([]byte("404: Not Found"))
}
