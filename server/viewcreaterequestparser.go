package server

import (
	"encoding/json"
	"fmt"
	"github.com/bradsk88/peertube-recommender/peertube"
	"github.com/pkg/errors"
	"net/http"
)

func NewViewCreateRequestParser() *ViewCreateRequestParser {
	return &ViewCreateRequestParser{}
}

type ViewCreateRequestParser struct {
}

func (parser *ViewCreateRequestParser) Parse(r *http.Request) (*ViewCreateRequest, error) {
	// Take care to only return errors that are OK for users to see
	var req ViewCreateRequest
	d := json.NewDecoder(r.Body)
	err := d.Decode(&req)
	if err != nil {
		return nil, errors.Wrap(err, "Unable to parse request body as JSON")
	}
	if req.Origin.VideoID() == "" {
		return nil, fmt.Errorf("origin.videoId missing from request")
	}
	if req.VideoName == "" {
		return nil, fmt.Errorf("videoName missing from request")
	}
	if req.VideoID == "" {
		return nil, fmt.Errorf("videoId missing from request")
	}
	if req.VideoURI == "" {
		return nil, fmt.Errorf("videoUri missing from request")
	}
	if req.WatchSeconds == 0 {
		return nil, fmt.Errorf("watchSeconds must be an integer > 0")
	}
	return &req, nil
}

type ViewCreateRequest struct {
	Origin       peertube.SimpleVideoIdentification `json:"origin"`
	VideoURI     string                             `json:"videoUri"`
	VideoName    string                             `json:"videoName"`
	WatchSeconds int64                              `json:"watchSeconds"`
	VideoID      string                             `json:"videoId"`
}
