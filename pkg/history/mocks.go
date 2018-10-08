package history

import (
	"fmt"
)

//NewMockRepository creates a new MockRepository
func NewMockRepository() *MockRepository {
	return &MockRepository{
		entries:    []Immutable{},
		RaiseOnGet: false,
	}
}

type MockRepository struct {
	entries    []Immutable
	RaiseOnGet bool
}

func (m *MockRepository) List(videoID string) ([]Immutable, error) {
	if m.RaiseOnGet {
		return nil, fmt.Errorf("raising on LOOKUP (get) as requested")
	}
	var filtered []Immutable
	for _, e := range m.entries {
		if e.origin.VideoID() == videoID {
			filtered = append(filtered, e)
		}
	}
	return filtered, nil
}

func (m *MockRepository) AddHistory(history History) error {
	m.entries = append(m.entries, NewImmutable(history.UserID(), history.Origin(), history.Video(), history.WatchSeconds()))
	return nil
}
