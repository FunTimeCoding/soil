package backup

import (
	"github.com/funtimecoding/soil/pkg/errors/not_found"
	"github.com/funtimecoding/soil/pkg/tool/gopostgresd/constant"
	"os"
)

func (s *Store) Existing(
	instance string,
	name string,
) (string, error) {
	if _, e := ParseName(name); e != nil {
		return "", e
	}

	result := s.dumpPath(instance, name)

	if _, e := os.Stat(result); os.IsNotExist(e) {
		return "", not_found.New(constant.BackupKind, name)
	} else if e != nil {
		return "", e
	}

	return result, nil
}
