package recommendations

import "github.com/bradsk88/peertube-recommender/peertube"

type Recommender interface {
	List(origin peertube.VideoIdentification) ([]Recommendation, error)
}

type Recommendation interface {
	peertube.DestinationVideo
}

func AreEqual(r1 Recommendation, r2 Recommendation) bool {
	if r1.Name() != r2.Name() {
		return false
	}
	if r1.ID() != r2.ID() {
		return false
	}
	if r1.URI() != r2.URI() {
		return false
	}
	return true
}

type Repository interface {
}
