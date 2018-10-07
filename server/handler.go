package server

import (
	"fmt"
	"github.com/inconshreveable/log15"
	"github.com/pkg/errors"
	"net/http"
)

type handler struct {
	Logger log15.Logger
}

func (h *handler) error(w http.ResponseWriter, msg string, err error, code int) {
	stack := fmt.Sprintf("%s", errors.WithStack(err).Error())
	h.Logger.Error(msg, "Error", stack)
	w.WriteHeader(code)
}
