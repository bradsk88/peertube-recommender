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
	expected := []recommendations.Recommendation{}
	tests.AssertRecommendationsEqual(t, expected, l)

}
