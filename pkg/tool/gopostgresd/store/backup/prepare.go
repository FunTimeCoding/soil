package backup

import (
	"github.com/funtimecoding/soil/pkg/errors/not_configured"
	"github.com/funtimecoding/soil/pkg/tool/gopostgresd/constant"
	"os"
)

func (s *Store) Prepare(
	instance string,
	name string,
) (string, error) {
	if s.directory == "" {
		return "", not_configured.Format(constant.BackupDirectoryUnset)
	}

	if _, e := ParseName(name); e != nil {
		return "", e
	}

	if e := os.MkdirAll(
		s.instanceDirectory(instance),
		constant.BackupDirectoryMode,
	); e != nil {
		return "", e
	}

	return s.dumpPath(instance, name), nil
}
