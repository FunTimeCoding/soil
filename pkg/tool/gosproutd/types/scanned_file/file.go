package scanned_file

import "time"

type File struct {
	Name        string
	Path        string
	ContentHash string
	Content     string
	ModifiedAt  time.Time
}
