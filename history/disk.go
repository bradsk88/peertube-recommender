package history

import (
	"encoding/json"
	"fmt"
	"github.com/bradsk88/peertube-recommender/peertube"
	"github.com/peterbourgon/diskv"
	"github.com/pkg/errors"
	"io/ioutil"
	"os"
	"time"
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

func (d *DiskBackedRepository) List(videoID string) ([]Immutable, error) {
	path := d.videoEventsPath(videoID)
	files, err := ioutil.ReadDir(path)
	if err != nil {
		//if strings.Contains(err.Error(), "no such file or directory") {
		//	return []Immutable{}, nil
		//}
		return nil, errors.Wrap(err, "Failed to load history for listing recommendations.")
	}
	out, err := d.readFilesForList(path, files)
	if err != nil {
		return nil, errors.Wrap(err, "Failed to read history files")
	}
	return out, nil
}

func (d *DiskBackedRepository) readFilesForList(path string, files []os.FileInfo) ([]Immutable, error) {
	out := make([]Immutable, len(files))
	for i, f := range files {
		fp := fmt.Sprintf("%s/%s", path, f.Name())
		bytes, err := ioutil.ReadFile(fp)
		if err != nil {
			return nil, errors.Wrapf(err, "Failed to read file: %s", f.Name())
		}
		var o jsonHistory
		err = json.Unmarshal(bytes, &o)
		if err != nil {
			return nil, errors.Wrapf(err, "Failed to read file: %s", f.Name())
		}
		d := peertube.NewImmutableDestinationVideo(o.VideoID, o.URI, o.Name)
		out[i] = NewImmutable(o.User, peertube.NewVideoIdentifiers(o.Origin), d, o.WatchTime)
	}
	return out, nil
}

type jsonHistory struct {
	Origin    string `json:"origin"`
	User      string `json:"user"`
	Name      string `json:"name"`
	URI       string `json:"uri"`
	VideoID   string `json:"videoId"`
	WatchTime int64  `json:"watchTime"`
}

func (d *DiskBackedRepository) AddHistory(h History) error {
	t := time.Now()
	path := d.videoEventsPath(h.Origin().VideoID())
	err := os.MkdirAll(path, os.ModePerm)
	if err != nil {
		return errors.Wrap(err, "Failed to create directory")
	}
	jsh := jsonHistory{
		Origin:    h.Origin().VideoID(),
		User:      h.UserID(),
		WatchTime: h.WatchSeconds(), // TODO: Average with prev history
		Name:      h.Video().Name(),
		VideoID:   h.Video().ID(),
		URI:       h.Video().URI(),
	}
	s, err := json.Marshal(jsh)
	if err != nil {
		return errors.Wrap(err, "Failed to serialize history for storage")
	}

	uuid := "abc" // TODO: Get an actual UUID
	filename := fmt.Sprintf("%s/%s-%s", path, t.Format("20060102150405"), uuid)
	err = ioutil.WriteFile(filename, s, 0644)
	err = d.store.Write(h.Origin().VideoID(), s)
	if err != nil {
		return errors.Wrap(err, "Failed to write to disk")
	}
	return nil
}

func (d *DiskBackedRepository) videoEventsPath(originVideoId string) string {
	return fmt.Sprintf("db/events/videos/%s", originVideoId)
}
