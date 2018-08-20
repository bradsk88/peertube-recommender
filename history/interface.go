package history

type History interface {
	UserID() string
	Video() Destination
	WatchSeconds() int64
}
