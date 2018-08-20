package justintimehistory

import (
	"fmt"
	"github.com/bradsk88/peertube-recommender/history"
	"github.com/bradsk88/peertube-recommender/peertube"
	"github.com/bradsk88/peertube-recommender/recommendations"
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
	//history, err := r.historyRepo.List(origin.VideoID())
	return nil, fmt.Errorf("not implemented")
}
