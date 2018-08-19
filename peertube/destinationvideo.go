package peertube

type DestinationVideo interface {
	URI() string
	Name() string
}
