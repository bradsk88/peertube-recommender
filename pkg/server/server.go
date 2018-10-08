package server

import (
	"github.com/bradsk88/peertube-recommender/pkg/history"
	"github.com/bradsk88/peertube-recommender/pkg/recommendations"
	"github.com/inconshreveable/log15"
	"github.com/pkg/errors"
	"net/http"
)

type HTTP struct {
	Logger      log15.Logger
	HistoryRepo history.Repository
	Recommender recommendations.Recommender
}

func (s *HTTP) Serve() error {
	mux := http.ServeMux{}
	mux.Handle("/recommendations", NewListHandler(s.Logger, s.Recommender))
	mux.Handle("/view", NewViewCreateHandler(s.Logger, s.HistoryRepo))

	err := http.ListenAndServe(":9999", &mux)
	if err != nil {
		return errors.Wrap(err, "Unable to serve")
	}
	return nil
}
