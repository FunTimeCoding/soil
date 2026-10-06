package index

import "time"

type fileSum struct {
	size     int64
	modified time.Time
	sum      [32]byte
}
