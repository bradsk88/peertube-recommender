package experimental

import (
	"github.com/bradsk88/peertube-recommender/history"
	"github.com/bradsk88/peertube-recommender/peertube"
	"github.com/bradsk88/peertube-recommender/recommendations"
	"github.com/bradsk88/peertube-recommender/videorepo"
	"github.com/pkg/errors"
	"testing"
)

func assertRecommendationsEqual(t *testing.T, expected []recommendations.Recommendation, actual []recommendations.Recommendation) {
	failed := false
	if expected == nil && actual != nil {
		failed = true
	}
	if !failed && expected != nil && actual == nil {
		failed = true
	}
	if !failed && len(expected) != len(actual) {
		failed = true
	}
	if !failed {
		for i, val := range expected {
			if val != actual[i] {
				failed = true
			}
		}
	}
	if failed {
		t.Errorf("Recommendations did not match expectation:\n"+
			"Expected:\n"+
			"%s\n"+
			"but got:\n"+
			"%s",
			expected, actual,
		)
	}
}

func TestShouldReturnErrorIfVideoRepositoryFailsToProvideVideos(t *testing.T) {
	repo := videorepo.NewMockRepository()
	repo.RaiseOnGet = true
	r := NewRecommender(repo, history.NewMockRepository())
	origin := peertube.NewVideoIdentifiers("origin video ID")
	_, err := r.List(origin)
	if err == nil {
		t.Errorf("Should have raised an error")
	}
}

func TestShouldReturnErrorIfHistoryRepositoryFailsToProvideVideos(t *testing.T) {
	histRepo := history.NewMockRepository()
	histRepo.RaiseOnGet = true
	r := NewRecommender(videorepo.NewMockRepository(), histRepo)
	origin := peertube.NewVideoIdentifiers("origin video ID")
	_, err := r.List(origin)
	if err == nil {
		t.Errorf("Should have raised an error")
	}
}

func TestShouldReturnARecommendation(t *testing.T) {
	repo := videorepo.NewMockRepository()
	repo.AddVideo(peertube.NewImmutableDestinationVideo("VidID", "VidName"))
	noHistory := history.NewMockRepository()
	r := NewRecommender(repo, noHistory)
	origin := peertube.NewVideoIdentifiers("origin video ID")
	results, err := r.List(origin)
	if err != nil {
		t.Errorf("Failed to list recommendations: %s", errors.WithStack(err))
		return
	}
	if results == nil {
		t.Errorf("Should have returned a recommendation, got nil")
		return
	}
	if len(results) == 0 {
		t.Errorf("Should have returned a recommendation, got []")
		return
	}
}

func TestShouldPreferAFullViewOverAPartialView(t *testing.T) {
	videos := videorepo.NewMockRepository()
	vidHistory := history.NewMockRepository()
	r := NewRecommender(videos, vidHistory)

	origin := peertube.NewVideoIdentifiers("origin video ID")

	destination1 := peertube.NewImmutableDestinationVideo("http://example.com/1", "Video 1")
	watchPercent1 := 1.0
	vidHistory.AddHistory(history.NewImmutable(origin, destination1, watchPercent1))

	destination2 := peertube.NewImmutableDestinationVideo("http://example.com/2", "Video 2")
	watchPercent2 := 0.99
	vidHistory.AddHistory(history.NewImmutable(origin, destination2, watchPercent2))

	expected := []recommendations.Recommendation{
		recommendations.NewImmutable(destination1.Name(), destination1.URI()),
		recommendations.NewImmutable(destination2.Name(), destination2.URI()),
	}

	results, _ := r.List(origin)
	assertRecommendationsEqual(t, expected, results)
}

func TestShouldPreferHistoryOverOtherVideos(t *testing.T) {
	videos := videorepo.NewMockRepository()
	vidHistory := history.NewMockRepository()
	r := NewRecommender(videos, vidHistory)

	origin := peertube.NewVideoIdentifiers("origin video ID")

	randomVideo := peertube.NewImmutableDestinationVideo(
		"http://example.com/random-video",
		"Random Video",
	)
	videos.AddVideo(randomVideo)

	previousSuccess := peertube.NewImmutableDestinationVideo(
		"http://example.com/previous-success",
		"Previous Successful Recommendation",
	)
	h := history.NewImmutable(origin, previousSuccess, 0.5)
	vidHistory.AddHistory(h)

	results, _ := r.List(origin)

	expected := []recommendations.Recommendation{
		recommendations.NewImmutable(previousSuccess.Name(), previousSuccess.URI()),
		recommendations.NewImmutable(randomVideo.Name(), randomVideo.URI()),
	}

	assertRecommendationsEqual(t, expected, results)
}
