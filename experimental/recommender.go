package experimental

import (
	"github.com/bradsk88/peertube-recommender/history"
	"github.com/bradsk88/peertube-recommender/peertube"
	"github.com/bradsk88/peertube-recommender/recommendations"
	"github.com/bradsk88/peertube-recommender/videorepo"
	"github.com/pkg/errors"
)

func NewRecommender(videoRepo videorepo.Repository, historyRepo history.Repository) *Recommender {
	return &Recommender{
		videos:  videoRepo,
		history: historyRepo,
	}
}

type Recommender struct {
	videos  videorepo.Repository
	history history.Repository
}

func (r *Recommender) List(origin peertube.VideoIdentification) ([]recommendations.Recommendation, error) {
	h, err := r.history.LookupForOrigin(origin)
	if err != nil {
		return nil, errors.Wrapf(err, "Failed to look up history for origin: %s", origin.VideoID())
	}

	videos, err := r.videos.GetNewVideos()
	if err != nil {
		return nil, errors.Wrap(err, "Failed to get videos from repo")
	}

	hs := destinationToRecommendations(historyToDestination(h))
	randoms := destinationToRecommendations(videos)
	return append(hs, randoms...), nil
}

func historyToDestination(h []history.History) []peertube.DestinationVideo {
	var out = make([]peertube.DestinationVideo, len(h))
	for i, hs := range h {
		out[i] = hs.Video()
	}
	return out
}

func destinationToRecommendations(v []peertube.DestinationVideo) []recommendations.Recommendation {
	var out = make([]recommendations.Recommendation, len(v))
	for i, d := range v {
		out[i] = recommendations.NewImmutable(d.ID(), d.Name(), d.URI())
	}
	return out
}
