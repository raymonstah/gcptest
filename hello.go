package main

import (
	"fmt"
	"net/http"
	"strings"
)

func GetHello(w http.ResponseWriter, r *http.Request) {
	name := r.URL.RequestURI()
	name = strings.Trim(name, "/")

	helloStr := fmt.Sprintf("Hello, %v", name)
	_, err := w.Write([]byte(helloStr))
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		panic(fmt.Errorf("error writing to writer: %w", err).Error())
	}

}
