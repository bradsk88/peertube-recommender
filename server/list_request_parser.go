package server

import (
	"github.com/bradsk88/peertube-recommender/peertube"
	"net/http"
)

type ListRequestParser struct {
}

type ListRequest struct {
	Origin peertube.VideoIdentification
}

func (parser *ListRequestParser) parse(request *http.Request) (*ListRequest, error) {
	origin := peertube.NewVideoIdentifiers("TODO") // TODO: Get this from the request
	return &ListRequest{
		Origin: origin,
	}, nil
}
