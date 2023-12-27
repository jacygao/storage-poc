package storage

type Item struct {
	ID           string `json:"id"`
	FileName     string `json:"filename"`
	Path         string `json:"path"`
	Query        string `json:"query"`
	CreationDate string `json:"createDate"`
	Version      string `json:"version"`
}

type Storage interface {
}
