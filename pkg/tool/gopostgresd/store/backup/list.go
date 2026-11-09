package backup

import (
	"github.com/funtimecoding/soil/pkg/tool/gopostgresd/constant"
	"github.com/funtimecoding/soil/pkg/tool/gopostgresd/types/backup"
	"os"
	"slices"
	"strings"
)

func (s *Store) List(instance string) ([]*backup.Backup, error) {
	entries, e := os.ReadDir(s.instanceDirectory(instance))

	if os.IsNotExist(e) {
		return []*backup.Backup{}, nil
	}

	if e != nil {
		return nil, e
	}

	result := []*backup.Backup{}

	for _, entry := range entries {
		name, isDump := strings.CutSuffix(
			entry.Name(),
			constant.BackupExtension,
		)

		if !isDump {
			continue
		}

		at, f := ParseName(name)

		if f != nil {
			continue
		}

		i, g := entry.Info()

		if g != nil {
			return nil, g
		}

		result = append(result, backup.New(instance, name, at, i.Size()))
	}

	slices.SortFunc(
		result,
		func(x *backup.Backup, y *backup.Backup) int {
			return y.Time.Compare(x.Time)
		},
	)

	return result, nil
}
