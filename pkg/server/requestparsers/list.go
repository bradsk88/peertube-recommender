package requestparsers

import (
	"encoding/json"
	"fmt"
	"github.com/bradsk88/peertube-recommender/pkg/serverrver/requests"
	"net/http"
)

func ForList() *List {
	return &List{}
}

type List struct {
}

func (parser *List) Parse(request *http.Request) (*requests.ListRequest, error) {
	// Take care to only return errors that are OK for users to see
	var req requests.ListRequest
	d := json.NewDecoder(request.Body)
	err := d.Decode(&req)
	if err != nil {
		return nil, fmt.Errorf("unable to parse request body. (JSON Error: %s)", err.Error())
	}
	if req.Origin.VideoID() == "" {
		return nil, fmt.Errorf("origin.videoId missing from request")
	}
	return &req, nil
}
