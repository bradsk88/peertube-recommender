package history

import "github.com/bradsk88/peertube-recommender/peertube"

func NewImmutable(userId string, origin peertube.VideoIdentification,
	dest Destination, watchSeconds int64) Immutable {
	return Immutable{
		origin:       origin,
		destination:  dest,
		watchSeconds: watchSeconds,
		userId:       userId,
	}
}

type Immutable struct {
	origin       Origin
	destination  Destination
	userId       string
	watchSeconds int64
}

func (i Immutable) UserID() string {
	return i.userId
}

func (i Immutable) WatchSeconds() int64 {
	return i.watchSeconds
}

func (i Immutable) Video() Destination {
	return i.destination
}
