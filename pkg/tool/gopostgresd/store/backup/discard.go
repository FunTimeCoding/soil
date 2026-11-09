package backup

import (
	"errors"
	"github.com/funtimecoding/soil/pkg/strings/join"
	"github.com/funtimecoding/soil/pkg/tool/gopostgresd/constant"
	"os"
)

func (s *Store) Discard(path string) error {
	var result []error

	for _, p := range []string{
		path,
		join.Empty(path, constant.ChecksumExtension),
	} {
		if e := os.Remove(p); e != nil && !os.IsNotExist(e) {
			result = append(result, e)
		}
	}

	return errors.Join(result...)
}
