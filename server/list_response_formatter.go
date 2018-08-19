package server

import (
	"encoding/json"
	"github.com/bradsk88/peertube-recommender/peertube"
	"github.com/bradsk88/peertube-recommender/recommendations"
	"github.com/pkg/errors"
)

type recommendationResponse struct {
	Origin          originData                  `json:"origin"`
	Recommendations []recommendations.Immutable `json:"recommendations"`
}

type originData struct {
	ID string `json:"video_id"`
}

type ListResponseFormatter struct {
}

type origin peertube.VideoIdentification

func (l *ListResponseFormatter) format(o origin, r []recommendations.Recommendation) ([]byte, error) {
	d := normalizeData(r)
	rr := recommendationResponse{
		Origin: originData{
			ID: o.VideoID(),
		},
		Recommendations: d,
	}
	s, err := json.Marshal(rr)
	if err != nil {
		return nil, errors.Wrap(err, "Unable to format recommendation")
	}
	return s, nil
}

func normalizeData(r []recommendations.Recommendation) []recommendations.Immutable {
	out := make([]recommendations.Immutable, len(r))
	for i, rec := range r {
		out[i] = recommendations.NewImmutable(rec.Name(), rec.URI())
	}
	return out
}
