package tree_entry

import "time"

type Entry struct {
	Size      int64
	Modified  time.Time
	Directory bool
	Linked    bool
}
