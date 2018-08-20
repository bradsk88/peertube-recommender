package peertube

func NewImmutableDestinationVideo(videoId string, uri string, name string) ImmutableDestinationVideo {
	return ImmutableDestinationVideo{
		uri:     uri,
		name:    name,
		videoId: videoId,
	}
}

type ImmutableDestinationVideo struct {
	name    string
	uri     string
	videoId string
}

func (v ImmutableDestinationVideo) ID() string {
	return v.videoId
}

func (v ImmutableDestinationVideo) URI() string {
	return v.uri
}

func (v ImmutableDestinationVideo) Name() string {
	return v.name
}
