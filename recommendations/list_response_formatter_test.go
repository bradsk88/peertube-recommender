package recommendations

import (
	"github.com/bradsk88/peertube-recommender/peertube"
	"testing"
)

func TestFormatGeneratesCorrectResponseForSimpleRecommendations(t *testing.T) {
	formatter := &ListResponseFormatter{}
	recommendations := []Recommendation{
		recommendation{
			NameValue: "Name1",
			URIValue:  "URL1",
		},
		recommendation{
			NameValue: "Name2",
			URIValue:  "URL2",
		},
	}
	origin := peertube.NewVideoIdentifiers("origin1")
	actual, err := formatter.format(origin, recommendations)
	if err != nil {
		t.Fail()
	}
	expected := `{"origin":{"video_id":"origin1"},"recommendations":[{"name":"Name1","uri":"URL1"},{"name":"Name2","uri":"URL2"}]}`
	if string(actual) != expected {
		t.Errorf("Expected:\n%s\nGot:\n%s", expected, string(actual))
	}
}
