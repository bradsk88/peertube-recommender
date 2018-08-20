package server

import (
	"github.com/bradsk88/peertube-recommender/recommendations"
	log "github.com/inconshreveable/log15"
	"net/http"
)

func NewListHandler(l log.Logger, r recommendations.Recommender) *ListHandler {
	listHandler := &ListHandler{
		recommender: r,
		parser:      &ListRequestParser{},
		formatter:   &ListResponseFormatter{},
	}
	listHandler.Logger = l.New("handler", "recommendations:GET")
	return listHandler
}

type ListHandler struct {
	handler
	recommender recommendations.Recommender
	formatter   *ListResponseFormatter
	parser      *ListRequestParser
}

func (h *ListHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	req, err := h.parser.parse(r)
	if err != nil {
		h.error(w, "Error parsing list request", err, 500)
		return
	}
	rs, err := h.recommender.List(req.Origin)
	if err != nil {
		h.error(w, "Error getting recommendations:", err, 500)
		return
	}
	body, err := h.formatter.format(req.Origin, rs)
	if err != nil {
		h.error(w, "Error formatting response:", err, 500)
		return
	}
	w.WriteHeader(200)
	w.Write(body)
}
