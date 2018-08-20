package history

type History interface {
	Video() Destination
	WatchSeconds() int64
}
