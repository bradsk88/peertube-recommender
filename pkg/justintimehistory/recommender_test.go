package justintimehistory

import (
	"github.com/bradsk88/peertube-recommender/pkg/history"
	"github.com/bradsk88/peertube-recommender/pkg/peertube"
	"github.com/bradsk88/peertube-recommender/pkg/recommendations"
	"github.com/bradsk88/peertube-recommender/pkg/tests"
	"testing"
)

func TestShouldReturnEmptyListIfNoHistoryExists(t *testing.T) {
	r := Recommender{
		historyRepo: history.NewMockRepository(),
	}
	l, err := r.List(peertube.NewVideoIdentifiers("vid"))
	if err != nil {
		t.Errorf("Should not have returned error: %s", err.Error())
		return
	}
	var expected []recommendations.Recommendation
	err = tests.AreRecommendationsEqual(expected, l)
	if err != nil {
		t.Error(err.Error())
	}
}

func TestShouldReturnCorrectRecommendations(t *testing.T) {
	repo := history.NewMockRepository()
	origin := peertube.NewVideoIdentifiers("vid1")
	dest1 := peertube.NewImmutableDestinationVideo("vid2", "http://example.com/2", "Video 2")
	dest2 := peertube.NewImmutableDestinationVideo("vid3", "http://example.com/3", "Video 3")
	repo.AddHistory(history.NewImmutable("user1", origin, dest1, 100))
	repo.AddHistory(history.NewImmutable("user2", origin, dest2, 100))
	r := Recommender{
		historyRepo: repo,
	}
	l, err := r.List(origin)
	if err != nil {
		t.Errorf("Should not have returned error: %s", err.Error())
	}
	expected := []recommendations.Recommendation{
		dest1,
		dest2,
	}
	err = tests.AreRecommendationsEqual(expected, l)
	if err != nil {
		t.Errorf(err.Error())
	}
}

func TestShouldNotReturnUnrelatedVideos(t *testing.T) {
	repo := history.NewMockRepository()
	origin := peertube.NewVideoIdentifiers("vid1")
	differentOrigin := peertube.NewVideoIdentifiers("vid99")
	dest1 := peertube.NewImmutableDestinationVideo("vid2", "http://example.com/2", "Video 2")
	dest2 := peertube.NewImmutableDestinationVideo("vid3", "http://example.com/3", "Video 3")
	repo.AddHistory(history.NewImmutable("user1", origin, dest1, 100))
	repo.AddHistory(history.NewImmutable("user2", differentOrigin, dest2, 100))
	r := Recommender{
		historyRepo: repo,
	}
	l, err := r.List(origin)
	if err != nil {
		t.Errorf("Should not have returned error: %s", err.Error())
	}
	expected := []recommendations.Recommendation{
		dest1,
	}
	err = tests.AreRecommendationsEqual(expected, l)
	if err != nil {
		t.Errorf(err.Error())
	}
}
