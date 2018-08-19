package history

type Repository interface {
	LookupForOrigin(origin Origin) ([]History, error)
}
