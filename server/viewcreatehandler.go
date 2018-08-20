package server

import (
	"fmt"
	"github.com/bradsk88/peertube-recommender/history"
	"github.com/bradsk88/peertube-recommender/peertube"
	"github.com/inconshreveable/log15"
	"net/http"
)

func NewViewCreateHandler(l log15.Logger, repository history.Repository) *ViewCreateHandler {
	createHandler := &ViewCreateHandler{
		parser: NewViewCreateRequestParser(),
		repo:   repository,
	}
	createHandler.Logger = l.New("handler", "view:PUT")
	return createHandler
}

type ViewCreateHandler struct {
	handler
	parser *ViewCreateRequestParser
	repo   history.Repository
}

func (h *ViewCreateHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	req, err := h.parser.Parse(r)
	if err != nil {
		h.error(w, "Error parsing create request", err, 400)
		w.Write([]byte(fmt.Sprintf("Invalid request: %s", err.Error())))
		return
	}
	d := peertube.NewImmutableDestinationVideo(req.VideoID, req.VideoURI, req.VideoName)
	err = h.repo.AddHistory(history.NewImmutable(&req.Origin, d, req.WatchSeconds))
	if err != nil {
		h.error(w, "Error storing view:", err, 500)
		return
	}
	w.WriteHeader(201)
}
