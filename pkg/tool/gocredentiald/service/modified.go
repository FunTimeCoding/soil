package service

import (
	"os"
	"time"
)

func modified(path string) time.Time {
	if i, e := os.Stat(path); e == nil {
		return i.ModTime()
	}

	return time.Time{}
}
