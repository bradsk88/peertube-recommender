package justintimehistory

import (
	"github.com/bradsk88/peertube-recommender/pkg/history"
	"github.com/bradsk88/peertube-recommender/pkg/peertube"
	"github.com/bradsk88/peertube-recommender/pkg/recommendations"
	"github.com/pkg/errors"
)

func NewRecommender(historyRepo history.Repository) *Recommender {
	return &Recommender{
		historyRepo: historyRepo,
	}
}

// Recommender reads the history repo just-in-time to generate recommendations
type Recommender struct {
	historyRepo history.Repository
}

func (r *Recommender) List(origin peertube.VideoIdentification) ([]recommendations.Recommendation, error) {
	histories, err := r.historyRepo.List(origin.VideoID())
	if err != nil {
		return nil, errors.Wrapf(err, "Failed to load history for origin video: %s", origin.VideoID())
	}
	var output []recommendations.Recommendation
	for _, hItem := range histories {
		v := recommendations.NewImmutable(hItem.Video().ID(), hItem.Video().Name(), hItem.Video().URI())
		output = append(output, v)
	}
	return output, nil
}
