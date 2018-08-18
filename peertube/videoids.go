package peertube

func NewVideoIdentifiers(videoID string) VideoIdentification {
	return &SimpleVideoIdentification{
		ID: videoID,
	}
}

type SimpleVideoIdentification struct {
	ID string `json:"video_id"`
}

func (s *SimpleVideoIdentification) VideoID() string {
	return s.ID
}

type VideoIdentification interface {
	VideoID() string
}
