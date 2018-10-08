package requestparsers_test

import (
	"bytes"
	"github.com/bradsk88/peertube-recommender/pkg/serverrver/requestparsers"
	"io/ioutil"
	"net/http"
	"testing"
)

func fakeRequest(body string) *http.Request {
	return &http.Request{
		Body: ioutil.NopCloser(bytes.NewBufferString(body)),
	}
}

func TestShouldReturnErrorIfRequestIsEmpty(t *testing.T) {
	p := requestparsers.ForCreateVideoView()
	body := `{}`
	r := fakeRequest(body)
	_, err := p.Parse(r)
	if err == nil {
		t.Fail()
	}
}

func TestShouldNotReturnErrorIfAllValuesPresentInRequest(t *testing.T) {
	p := requestparsers.ForCreateVideoView()
	body := `{
		"origin": {
			"videoId": "V0"
		},
		"videoUri": "http://example.com",
		"videoName": "Day at the Zoo",
		"watchSeconds": 100,
		"videoId": "v1"
	}`
	r := fakeRequest(body)
	_, err := p.Parse(r)
	if err != nil {
		t.Errorf("Should not have returned an error: %s", err.Error())
	}
}

type missingValueTest struct {
	name string
	body string
}

func TestShouldReturnErrorIfMissingValue(t *testing.T) {
	p := requestparsers.ForCreateVideoView()
	tests := []missingValueTest{
		{
			name: "origin",
			body: `{
				"videoUri": "http://example.com",
				"videoName": "Day at the Zoo",
				"watchSeconds": 100,
				"videoId": "v1"
			}`,
		},
		{
			name: "videoUri",
			body: `{
				"origin": {
					"videoId": "V0"
				},
				"videoName": "Day at the Zoo",
				"watchSeconds": 100,
				"videoId": "v1"
			}`,
		},
		{
			name: "videoName",
			body: `{
				"origin": {
					"videoId": "V0"
				},
				"videoUri": "http://example.com",
				"watchSeconds": 100,
				"videoId": "v1"
			}`,
		},
		{
			name: "watchSeconds",
			body: `{
				"origin": {
					"videoId": "V0"
				},
				"videoUri": "http://example.com",
				"videoName": "Day at the Zoo",
				"videoId": "v1"
			}`,
		},
		{
			name: "videoId",
			body: `{
				"origin": {
					"videoId": "V0"
				},
				"videoUri": "http://example.com",
				"videoName": "Day at the Zoo",
				"watchSeconds": 100,
			}`,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(tt *testing.T) {
			r := fakeRequest(test.body)
			_, err := p.Parse(r)
			if err == nil {
				t.Errorf("Should have returned an error")
			}
		})
	}
}
