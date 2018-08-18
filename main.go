package main

import (
	"fmt"
	"github.com/bradsk88/peertube-recommender/recommendations"
	"github.com/inconshreveable/log15"
	"net/http"
	"os"
)

func main() {
	l := log15.New("app", "recommender")
	l.SetHandler(log15.StreamHandler(os.Stderr, log15.TerminalFormat()))

	mux := http.ServeMux{}
	mux.Handle("/recommendations", recommendations.NewHandler(l, recommendations.NewRecommenderMock()))

	err := http.ListenAndServe(":9999", &mux)
	panic(fmt.Sprintf("Failed to serve: %s", err.Error()))
}
