package brave

import (
	"github.com/funtimecoding/soil/pkg/brave/helper"
	"github.com/funtimecoding/soil/pkg/brave/preference"
	"github.com/funtimecoding/soil/pkg/brave/profile"
	"os"
)

func Profiles() ([]*profile.Profile, error) {
	p := helper.SettingsPath()
	entries, e := os.ReadDir(p)

	if e != nil {
		return nil, e
	}

	var result []*profile.Profile

	for _, f := range entries {
		if !profile.IsProfile(f) {
			continue
		}

		r, g := preference.Parse(f.Name())

		if g != nil {
			return nil, g
		}

		if r.Cookie.Accounts == "" {
			continue
		}

		parsed, h := profile.Parse(p, f.Name(), r.Cookie.Accounts)

		if h != nil {
			return nil, h
		}

		result = append(result, parsed)
	}

	return result, nil
}
