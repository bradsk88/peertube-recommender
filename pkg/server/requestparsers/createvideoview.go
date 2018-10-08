package requestparsers

import (
	"encoding/json"
	"fmt"
	"github.com/bradsk88/peertube-recommender/pkg/serverrver/requests"
	"github.com/pkg/errors"
	"net/http"
)

func ForCreateVideoView() *CreateVideoView {
	return &CreateVideoView{}
}

type CreateVideoView struct {
}

func (parser *CreateVideoView) Parse(r *http.Request) (*requests.CreateVideoViewRequest, error) {
	// Take care to only return errors that are OK for users to see
	var req requests.CreateVideoViewRequest
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
