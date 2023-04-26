package handlers

import (
	"fmt"
	"net/http"
)

func NewHelloHandlers() *HelloHandlers {
	return &HelloHandlers{}
}

type HelloHandlers struct{}

func (hh *HelloHandlers) GetHandlers() Handlers {
	return map[string]func(w http.ResponseWriter, r *http.Request){
		"/paul":   Hello("Paul"),
		"/bertie": Hello("Bertie"),
	}
}

func Hello(name string) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(fmt.Sprintf("Hello %s!", name)))
	}
}
