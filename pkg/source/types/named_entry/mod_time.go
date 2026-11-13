package named_entry

import "time"

func (e *Entry) ModTime() time.Time {
	return e.entry.Modified
}
