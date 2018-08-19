package peertube

func NewImmutableDestinationVideo(uri string, name string) ImmutableDestinationVideo {
	return ImmutableDestinationVideo{
		uri:  uri,
		name: name,
	}
}

type ImmutableDestinationVideo struct {
	name string
	uri  string
}

func (v ImmutableDestinationVideo) URI() string {
	return v.uri
}

func (v ImmutableDestinationVideo) Name() string {
	return v.name
}
