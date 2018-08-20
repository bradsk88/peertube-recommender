package history

type Repository interface {
	AddHistory(h History) error
	List(videoID string) []History
}
