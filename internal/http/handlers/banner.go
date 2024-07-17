package handlers

import (
	"fmt"
	"net/http"
)

func Banner(writer http.ResponseWriter, request *http.Request) {
	_, _ = fmt.Fprintf(writer, "usertwist\n")
}
