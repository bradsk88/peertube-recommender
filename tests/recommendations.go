package tests

import (
	"github.com/bradsk88/peertube-recommender/recommendations"
	"testing"
)

func AssertRecommendationsEqual(t *testing.T, expected []recommendations.Recommendation, actual []recommendations.Recommendation) {
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
