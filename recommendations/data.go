package recommendations

type recommendation struct {
	NameValue string `json:"name"`
	URIValue  string `json:"uri"`
}

func (r recommendation) Name() string {
	return r.NameValue
}

func (r recommendation) URI() string {
	return r.URIValue
}
