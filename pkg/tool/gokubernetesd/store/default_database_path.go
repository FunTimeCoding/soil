package store

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/system"
	"github.com/funtimecoding/soil/pkg/tool/gokubernetesd/constant"
	"path/filepath"
)

func DefaultDatabasePath() string {
	n := constant.Identity.Name()
	d := filepath.Join(system.Home(), ".local", "share", n)
	system.MakeDirectory(d)

	return filepath.Join(d, fmt.Sprintf("%s.sqlite", n))
}
