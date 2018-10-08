package history

import "github.com/bradsk88/peertube-recommender/pkg/peertube"

type History interface {
	UserID() string
	Video() Destination
	WatchSeconds() int64
	Origin() peertube.VideoIdentification
}