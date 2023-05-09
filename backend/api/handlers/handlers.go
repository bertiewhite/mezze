package handlers

import "net/http"

type Handlers map[string]func(w http.ResponseWriter, r *http.Request)

type HandlerProvider interface {
	GetHandlers() Handlers
}
