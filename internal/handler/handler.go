package handler

import "net/http"

type Handler interface {
	Handle() http.Handler
	Pattern() string
}
