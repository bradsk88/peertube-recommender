package history

import "fmt"

//NewMockRepository creates a new MockRepository
func NewMockRepository() *MockRepository {
	return &MockRepository{
		entries:    []History{},
		RaiseOnGet: false,
	}
}

type MockRepository struct {
	entries    []History
	RaiseOnGet bool
}

func (m *MockRepository) List(videoID string) ([]History, error) {
	if m.RaiseOnGet {
		return nil, fmt.Errorf("raising on LOOKUP (get) as requested")
	}
	return m.entries, nil
}

func (m *MockRepository) AddHistory(history History) error {
	m.entries = append(m.entries, history)
	return nil
}
