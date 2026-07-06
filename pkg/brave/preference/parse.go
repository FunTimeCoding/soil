package preference

import (
	"github.com/funtimecoding/soil/pkg/brave/constant"
	"github.com/funtimecoding/soil/pkg/brave/helper"
	"github.com/funtimecoding/soil/pkg/notation"
	"github.com/funtimecoding/soil/pkg/system"
	"github.com/funtimecoding/soil/pkg/system/join"
)

func Parse(profile string) *Preference {
	var result Preference
	p := join.Absolute(helper.SettingsPath(), profile)
	b := system.ReadBytes(p, constant.PreferencesFile)
	notation.MustDecodeBytes(b, &result, false)

	return &result
}
