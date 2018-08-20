package recommendations

func NewImmutable(videoId string, name string, uri string) Immutable {
	return Immutable{
		NameValue: name,
		URIValue:  uri,
		IDValue:   videoId,
	}
}

type Immutable struct {
	NameValue string `json:"name"`
	URIValue  string `json:"uri"`
	IDValue   string `json:"id"`
}

func (r Immutable) ID() string {
	return r.IDValue
}

func (r Immutable) Name() string {
	return r.NameValue
}

func (r Immutable) URI() string {
	return r.URIValue
}
