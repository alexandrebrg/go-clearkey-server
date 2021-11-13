package httpHandler

import "net/http"

type HttpHandler interface {
	GetKey(w http.ResponseWriter, r *http.Request)
	PostKey(w http.ResponseWriter, r *http.Request)
}
