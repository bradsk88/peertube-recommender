package recommendations

import "github.com/bradsk88/peertube-recommender/peertube"

func NewRecommenderMock() *RecommenderMock {
	return &RecommenderMock{}
}

type RecommenderMock struct {
}

func (*RecommenderMock) List(origin peertube.VideoIdentification) ([]Recommendation, error) {
	return []Recommendation{
		NewImmutable("Video 1", "http://example.com/1"),
		NewImmutable("Video 2", "http://example.com/2"),
	}, nil
}
