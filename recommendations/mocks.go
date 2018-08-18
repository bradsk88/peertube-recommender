package recommendations

import "github.com/bradsk88/peertube-recommender/peertube"

func NewRecommenderMock() *RecommenderMock {
	return &RecommenderMock{}
}

type RecommenderMock struct {
}

func (*RecommenderMock) List(origin peertube.VideoIdentification) ([]Recommendation, error) {
	return []Recommendation{
		recommendation{
			NameValue: "Video 1",
			URIValue:  "http://example.com/1",
		},
		recommendation{
			NameValue: "Video 2",
			URIValue:  "http://example.com/2",
		},
	}, nil
}
