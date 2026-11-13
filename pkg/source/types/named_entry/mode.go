package named_entry

import "io/fs"

func (e *Entry) Mode() fs.FileMode {
	if e.entry.Directory {
		return fs.ModeDir
	}

	return 0
}
