package justintimehistory

import (
	"github.com/bradsk88/peertube-recommender/history"
	"github.com/bradsk88/peertube-recommender/peertube"
	"github.com/bradsk88/peertube-recommender/recommendations"
	"github.com/bradsk88/peertube-recommender/tests"
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
	tests.AssertRecommendationsEqual(t, expected, l)
}

func TestShouldReturnCorrectRecommendations(t *testing.T) {
	// Add history for User 1 Video 1
	// Add history for User 1 Video 2
	// Add history for User 2 Video 1
	// Add history for User 2 Video 3
	// Get recommendations for origin Video 1
	// Expect to get Video 2, Video 3
	t.Errorf("not implemented")
}
