package history

import (
	"encoding/json"
	"github.com/peterbourgon/diskv"
	"github.com/pkg/errors"
)

func NewDiskBackedRepository() *DiskBackedRepository {
	flatTransform := func(s string) []string { return []string{} }
	db := diskv.New(diskv.Options{
		BasePath:     "peertube-recos",
		Transform:    flatTransform,
		CacheSizeMax: 1024 * 1024,
	})
	return &DiskBackedRepository{
		store: db,
	}
}

type DiskBackedRepository struct {
	store *diskv.Diskv
}

type jsonHistory struct {
	User      string `json:"user"`
	Name      string `json:"name"`
	URI       string `json:"uri"`
	WatchTime int64  `json:"watchTime"`
}

func (d *DiskBackedRepository) AddHistory(h History) error {
	jsh := jsonHistory{
		User:      h.UserID(),
		Name:      h.Video().Name(),
		URI:       h.Video().URI(),
		WatchTime: h.WatchSeconds(), // TODO: Average with prev history
	}
	s, err := json.Marshal(jsh)
	if err != nil {
		return errors.Wrap(err, "Failed to serialize history for storage")
	}
	err = d.store.Write(h.Video().ID(), s)
	if err != nil {
		return errors.Wrap(err, "Failed to write to disk")
	}
	return nil
}
