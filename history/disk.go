package history

import (
	"github.com/peterbourgon/diskv"
	"fmt"
	"encoding/json"
	"github.com/pkg/errors"
)

func NewDiskBackedRepository() *DiskBackedRepository {
	flatTransform := func(s string) []string {return []string{} }
	db := diskv.New(diskv.Options{
		BasePath: "peertube-recos",
		Transform: flatTransform,
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
	Name string `json:"name"`
	URI string `json:"uri"`
	AverageWatchTime int64 `json:"watchTime"`
	Views int64 `json:"views"`
}

func (d *DiskBackedRepository) AddHistory(h History) (error) {
	jsh := jsonHistory{
		Name: h.Video().Name(),
		URI: h.Video().URI(),
		AverageWatchTime: h.WatchSeconds(), // TODO: Average with prev history
		Views: 1, // TODO: Look up previous history and increment this
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

func (d *DiskBackedRepository) LookupForOrigin(origin Origin) ([]History, error) {
	return nil, fmt.Errorf("not implemented")
}
