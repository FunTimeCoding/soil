package transcript

import (
	"github.com/funtimecoding/soil/pkg/generative/anthropic/claude/permission/allowlist"
	generativeConstant "github.com/funtimecoding/soil/pkg/generative/constant"
	"github.com/funtimecoding/soil/pkg/system"
	"path/filepath"
)

func projectRules(
	cache map[string][]*allowlist.Rule,
	project string,
) ([]*allowlist.Rule, error) {
	if rules, okay := cache[project]; okay {
		return rules, nil
	}

	rules, e := allowlist.ReadAll(
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
	)

	if e != nil {
		return nil, e
	}

	cache[project] = rules

	return rules, nil
}
