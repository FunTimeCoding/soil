package transcript

import (
	generativeConstant "github.com/funtimecoding/soil/pkg/generative/constant"
	"github.com/funtimecoding/soil/pkg/system"
	"path/filepath"
)

func settingsPaths(project string) []string {
	return []string{
		filepath.Join(
			project,
			generativeConstant.SettingsDirectory,
			generativeConstant.SettingsFile,
		),
		filepath.Join(
			project,
			generativeConstant.SettingsDirectory,
			generativeConstant.SharedSettingsFile,
		),
		filepath.Join(
			system.Home(),
			generativeConstant.SettingsDirectory,
			generativeConstant.SharedSettingsFile,
		),
	}
}
