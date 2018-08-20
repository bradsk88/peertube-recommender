package history

import "github.com/bradsk88/peertube-recommender/peertube"

func NewImmutable(o peertube.VideoIdentification, d Destination, watchSeconds int64) Immutable {
	return Immutable{
		origin:       o,
		destination:  d,
		watchSeconds: watchSeconds,
	}
}

type Immutable struct {
	origin       Origin
	destination  Destination
	watchSeconds int64
}

func (i Immutable) WatchSeconds() int64 {
	return i.watchSeconds
}

func (i Immutable) Video() Destination {
	return i.destination
}
