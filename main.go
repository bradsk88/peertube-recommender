package main

import (
	"fmt"
	"github.com/bradsk88/peertube-recommender/history"
	"github.com/bradsk88/peertube-recommender/justintimehistory"
	"github.com/bradsk88/peertube-recommender/server"
	"github.com/inconshreveable/log15"
	"os"
)

func main() {
	l := log15.New("app", "recommender")
	l.SetHandler(log15.StreamHandler(os.Stderr, log15.TerminalFormat()))

	h := history.NewDiskBackedRepository()
	r := justintimehistory.NewRecommender(h)

	s := &server.HTTP{Logger: l, HistoryRepo: h, Recommender: r}
	err := s.Serve()
	if err != nil {
		panic(fmt.Sprintf("Failed to serve: %s", err.Error()))
	}
}
