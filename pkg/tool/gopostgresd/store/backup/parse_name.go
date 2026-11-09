package backup

import (
	"github.com/funtimecoding/soil/pkg/errors/validation"
	"github.com/funtimecoding/soil/pkg/tool/gopostgresd/constant"
	"time"
)

func ParseName(name string) (time.Time, error) {
	result, e := time.Parse(constant.BackupNameLayout, name)

	if e != nil || result.Format(constant.BackupNameLayout) != name {
		return time.Time{}, validation.New(
			constant.BackupNameRefused,
			name,
			constant.BackupNameLayout,
		)
	}

	return result, nil
}
