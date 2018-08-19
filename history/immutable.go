package history

func NewImmutable(o Origin, d Destination, watchPercent float64) Immutable {
	return Immutable{
		origin:       o,
		destination:  d,
		watchPercent: watchPercent,
	}
}

type Immutable struct {
	origin       Origin
	destination  Destination
	watchPercent float64
}

func (i Immutable) Video() Destination {
	return i.destination
}
