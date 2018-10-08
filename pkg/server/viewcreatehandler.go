package server

import (
	"fmt"
	"github.com/bradsk88/peertube-recommender/pkg/history"
	"github.com/bradsk88/peertube-recommender/pkg/peertube"
	"github.com/bradsk88/peertube-recommender/pkg/serverrver/requestparsers"
	"github.com/inconshreveable/log15"
	"net/http"
)

func NewViewCreateHandler(l log15.Logger, repository history.Repository) *ViewCreateHandler {
	createHandler := &ViewCreateHandler{
		parser: requestparsers.ForCreateVideoView(),
		repo:   repository,
	}
	createHandler.Logger = l.New("handler", "view:PUT")
	return createHandler
}

type ViewCreateHandler struct {
	handler
	parser *requestparsers.CreateVideoView
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
	i := history.NewImmutable(req.UserID, &req.Origin, d, req.WatchSeconds)
	err = h.repo.AddHistory(i)
	if err != nil {
		h.error(w, "Error storing view:", err, 500)
		return
	}
	w.WriteHeader(201)
}
