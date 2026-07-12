package service

import "time"

type DiscoveredFile struct {
	Name        string
	Path        string
	ContentHash string
	Content     string
	ModifiedAt  time.Time
}
