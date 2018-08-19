package recommendations

import "github.com/bradsk88/peertube-recommender/peertube"

type Recommender interface {
	List(origin peertube.VideoIdentification) ([]Recommendation, error)
}

type Recommendation interface {
	peertube.DestinationVideo
}
