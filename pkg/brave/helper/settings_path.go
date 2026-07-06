package helper

import (
	"github.com/funtimecoding/soil/pkg/brave/constant"
	"github.com/funtimecoding/soil/pkg/system"
	"path/filepath"
)

func SettingsPath() string {
	return filepath.Join(system.Home(), constant.BraveSettings)
}
