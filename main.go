package main

import (
	"net/http"
	"fmt"
	"github.com/bradsk88/peertube-recommender/recommendations"
)

func main() {
	mux := http.ServeMux{}
	mux.Handle("/recommendations", recommendations.NewHandler())
	err := http.ListenAndServe(":9999", &mux)
	panic(fmt.Sprintf("Failed to serve: %s", err.Error()))
}