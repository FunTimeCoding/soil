package mock_entry

import "io/fs"

func (f *Entry) Mode() fs.FileMode {
	if f.directory {
		return fs.ModeDir | 0755
	}

	return 0644
}
