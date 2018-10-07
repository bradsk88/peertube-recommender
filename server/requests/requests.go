package requests

import "github.com/bradsk88/peertube-recommender/peertube"

type ListRequest struct {
	Origin peertube.SimpleVideoIdentification `json:"origin"`
}

type CreateVideoViewRequest struct {
	UserID       string                             `json:"userId"`
	Origin       peertube.SimpleVideoIdentification `json:"origin"`
	VideoURI     string                             `json:"videoUri"`
	VideoName    string                             `json:"videoName"`
	WatchSeconds int64                              `json:"watchSeconds"`
	VideoID      string                             `json:"videoId"`
}
