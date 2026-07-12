package watcher

import "time"

type scannedFile struct {
	name        string
	path        string
	contentHash string
	content     string
	modifiedAt  time.Time
}
