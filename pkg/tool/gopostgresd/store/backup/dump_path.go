package backup

import (
	"github.com/funtimecoding/soil/pkg/strings/join"
	"github.com/funtimecoding/soil/pkg/tool/gopostgresd/constant"
	"path/filepath"
)

func (s *Store) dumpPath(
	instance string,
	name string,
) string {
	return filepath.Join(
		s.instanceDirectory(instance),
		join.Empty(name, constant.BackupExtension),
	)
}
