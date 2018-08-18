package recommendations

import (
	"net/http"
)

func NewHandler() *Handler {
	return &Handler{}
}

type Handler struct {
}

func (*Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(501)
}

