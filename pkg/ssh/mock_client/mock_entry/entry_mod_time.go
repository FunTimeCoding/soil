package mock_entry

import "time"

func (f *Entry) ModTime() time.Time { return time.Time{} }
