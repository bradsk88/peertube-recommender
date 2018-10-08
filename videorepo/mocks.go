package videorepo

import (
	"fmt"
	"github.com/bradsk88/peertube-recommender/pkg/peertube"
)

func NewMockRepository() *MockRepository {
	return &MockRepository{
		videos:     []peertube.DestinationVideo{},
		RaiseOnGet: false,
	}
}

type MockRepository struct {
	videos     []peertube.DestinationVideo
	RaiseOnGet bool
}

func (m *MockRepository) GetNewVideos() ([]peertube.DestinationVideo, error) {
	if m.RaiseOnGet {
		return nil, fmt.Errorf("raising on GET as requested")
	}
	return m.videos, nil
}

func (m *MockRepository) GetVideo(videoID string) (peertube.DestinationVideo, error) {
	if m.RaiseOnGet {
		return nil, fmt.Errorf("raising on GET as requested")
	}
	return m.videos[0], nil
}
func (m *MockRepository) AddVideo(d peertube.DestinationVideo) {
	m.videos = append(m.videos, d)
}
