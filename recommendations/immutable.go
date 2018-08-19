package recommendations

func NewImmutable(name string, uri string) Immutable {
	return Immutable{
		NameValue: name,
		URIValue:  uri,
	}
}

type Immutable struct {
	NameValue string `json:"name"`
	URIValue  string `json:"uri"`
}

func (r Immutable) Name() string {
	return r.NameValue
}

func (r Immutable) URI() string {
	return r.URIValue
}
