package tests

import (
	"fmt"
	"github.com/bradsk88/peertube-recommender/recommendations"
)

func AreRecommendationsEqual(expected []recommendations.Recommendation, actual []recommendations.Recommendation) error {
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
			if !recommendations.AreEqual(val, actual[i]) {
				failed = true
			}
		}
	}
	if failed {
		return fmt.Errorf("Recommendations did not match expectation:\n"+
			"Expected:\n"+
			"%s\n"+
			"but got:\n"+
			"%s",
			expected, actual,
		)
	}
	return nil
}
