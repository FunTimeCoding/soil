package preference

import (
	"github.com/funtimecoding/soil/pkg/brave/constant"
	"github.com/funtimecoding/soil/pkg/brave/helper"
	"github.com/funtimecoding/soil/pkg/notation"
	"os"
	"path/filepath"
)

func Parse(profile string) (*Preference, error) {
	b, e := os.ReadFile(
		filepath.Join(helper.SettingsPath(), profile, constant.PreferencesFile),
	)

	if e != nil {
		return nil, e
	}

	var result Preference

	if f := notation.DecodeBytes(b, &result); f != nil {
		return nil, f
	}

	return &result, nil
}
