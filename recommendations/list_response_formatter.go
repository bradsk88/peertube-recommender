package recommendations

import (
	"encoding/json"
	"github.com/bradsk88/peertube-recommender/peertube"
	"github.com/pkg/errors"
)

type recommendationResponse struct {
	Origin          originData       `json:"origin"`
	Recommendations []recommendation `json:"recommendations"`
}

type originData struct {
	ID string `json:"video_id"`
}

type ListResponseFormatter struct {
}

type origin = peertube.VideoIdentification

func (l *ListResponseFormatter) format(o origin, r []Recommendation) ([]byte, error) {
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

func normalizeData(r []Recommendation) []recommendation {
	out := make([]recommendation, len(r))
	for i, rec := range r {
		out[i] = recommendation{
			NameValue: rec.Name(),
			URIValue:  rec.URI(),
		}
	}
	return out
}
