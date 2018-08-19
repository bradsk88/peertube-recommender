package videorepo

import "github.com/bradsk88/peertube-recommender/peertube"

type Repository interface {
	GetVideo(videoID string) (peertube.DestinationVideo, error)
	GetNewVideos() ([]peertube.DestinationVideo, error)
}
