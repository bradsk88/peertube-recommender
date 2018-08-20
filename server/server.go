package server

import (
	"github.com/bradsk88/peertube-recommender/experimental"
	"github.com/bradsk88/peertube-recommender/history"
	"github.com/bradsk88/peertube-recommender/videorepo"
	"github.com/inconshreveable/log15"
	"github.com/pkg/errors"
	"net/http"
)

type HTTP struct {
	Logger      log15.Logger
	VideoRepo   videorepo.Repository
	HistoryRepo history.Repository
}

func (s *HTTP) Serve() error {
	mux := http.ServeMux{}
	r := experimental.NewRecommender(s.VideoRepo, s.HistoryRepo)
	mux.Handle("/recommendations", NewListHandler(s.Logger, r))
	mux.Handle("/view", NewViewCreateHandler(s.Logger, s.HistoryRepo))

	err := http.ListenAndServe(":9999", &mux)
	if err != nil {
		return errors.Wrap(err, "Unable to serve")
	}
	return nil
}
