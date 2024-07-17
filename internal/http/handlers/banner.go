package handlers

import (
	"fmt"
	"net/http"
)

func Banner() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		_, _ = fmt.Fprintf(writer, "usertwist\n")
	}
}
