package main

import (
	"fmt"
	"github.com/bradsk88/peertube-recommender/server"
	"github.com/inconshreveable/log15"
	"net/http"
	"os"
	"github.com/bradsk88/peertube-recommender/experimental"
	"github.com/bradsk88/peertube-recommender/videorepo"
	"github.com/bradsk88/peertube-recommender/history"
)

func main() {
	l := log15.New("app", "recommender")
	l.SetHandler(log15.StreamHandler(os.Stderr, log15.TerminalFormat()))

	mux := http.ServeMux{}
	r := experimental.NewRecommender(videorepo.NewMockRepository(), history.NewMockRepository())
	mux.Handle("/recommendations", server.NewHandler(l, r))

	err := http.ListenAndServe(":9999", &mux)
	panic(fmt.Sprintf("Failed to serve: %s", err.Error()))
}
