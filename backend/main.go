package main

import (
	"fmt"
	"mezze/api/handlers"
	"mezze/recipes/dummy"
	"net/http"
)

func main() {
	mux := http.NewServeMux()

	dummyRecipeStore := &dummy.DummyStore{}
	h := []handlers.HandlerProvider{
		handlers.NewHelloHandlers(),
		handlers.NewRecipesHandlers(dummyRecipeStore),
	}

	for _, providers := range h {
		for name, handler := range providers.GetHandlers() {
			fmt.Println(fmt.Sprintf("Setting up handling for %s", name))
			mux.HandleFunc(name, handler)
		}
	}

	http.ListenAndServe(":8080", mux)
}
