package layer

import (
	"io/fs"
	"os"
)

func (l *Layer) Read(path string) ([]byte, error) {
	if l.removed[path] {
		return nil, fs.ErrNotExist
	}

	if b, found := l.written[path]; found {
		return b, nil
	}

	if l.parent != nil {
		return l.parent.Read(path)
	}

	return os.ReadFile(path)
}
