package server

import (
	"github.com/bradsk88/peertube-recommender/recommendations"
	"github.com/bradsk88/peertube-recommender/server/requestparsers"
	"github.com/bradsk88/peertube-recommender/server/responseformatters"
	log "github.com/inconshreveable/log15"
	"net/http"
)

func NewListHandler(l log.Logger, r recommendations.Recommender) *ListHandler {
	listHandler := &ListHandler{
		recommender: r,
		parser:      &requestparsers.List{},
		formatter:   &responseformatters.List{},
	}
	listHandler.Logger = l.New("handler", "recommendations:GET")
	return listHandler
}

type ListHandler struct {
	handler
	recommender recommendations.Recommender
	formatter   *responseformatters.List
	parser      *requestparsers.List
}

func (h *ListHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	req, err := h.parser.Parse(r)
	if err != nil {
		h.error(w, "Error parsing create request", err, 400)
		w.Write([]byte("Invalid request body"))
		return
	}
	var hr []recommendations.Recommendation
	hr, err = h.recommender.List(&req.Origin)
	if err != nil {
		h.error(w, "Error loading recommendations:", err, 500)
		return
	}
	w.WriteHeader(200)
	bytes, err := h.formatter.Format(&req.Origin, hr)
	if err != nil {
		h.error(w, "Error formatting recommendations response:", err, 500)
	}
	w.Write(bytes)
}
