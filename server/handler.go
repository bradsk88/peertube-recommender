package server

import (
	"github.com/bradsk88/peertube-recommender/recommendations"
	log "github.com/inconshreveable/log15"
	"github.com/pkg/errors"
	"net/http"
)

func NewHandler(l log.Logger, r recommendations.Recommender) *Handler {
	logger := l.New("module", "recommendations")
	return &Handler{
		logger:      logger,
		recommender: r,
		parser:      &ListRequestParser{},
		formatter:   &ListResponseFormatter{},
	}
}

type Handler struct {
	logger      log.Logger
	recommender recommendations.Recommender
	formatter   *ListResponseFormatter
	parser      *ListRequestParser
}

func (h *Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
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

func (h *Handler) error(w http.ResponseWriter, msg string, err error, code int) {
	h.logger.Error(msg, "Error", errors.WithStack(err))
	w.WriteHeader(code)
}
