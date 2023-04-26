package main

import (
	"mezze/api/handlers"
	"net/http"
)

func main() {
	mux := http.NewServeMux()

	h := []handlers.HandlerProvider{
		handlers.NewHelloHandlers(),
	}

	for _, providers := range h {
		for name, handler := range providers.GetHandlers() {
			mux.HandleFunc(name, handler)
		}
	}

	http.ListenAndServe(":8080", mux)
}
