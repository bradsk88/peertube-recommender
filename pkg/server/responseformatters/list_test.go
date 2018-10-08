package responseformatters

import (
	"github.com/bradsk88/peertube-recommender/pkg/peertube"
	"github.com/bradsk88/peertube-recommender/pkg/recommendations"
	"testing"
)

func TestFormatGeneratesCorrectResponseForSimpleRecommendations(t *testing.T) {
	formatter := &List{}
	rs := []recommendations.Recommendation{
		recommendations.NewImmutable("V1", "Name1", "URL1"),
		recommendations.NewImmutable("V2", "Name2", "URL2"),
	}
	origin := peertube.NewVideoIdentifiers("origin1")
	actual, err := formatter.Format(origin, rs)
	if err != nil {
		t.Fail()
	}
	expected := `{"origin":{"videoId":"origin1"},"recommendations":[{"name":"Name1","uri":"URL1","id":"V1"},{"name":"Name2","uri":"URL2","id":"V2"}]}`
	if string(actual) != expected {
		t.Errorf("Expected:\n%s\nGot:\n%s", expected, string(actual))
	}
}
