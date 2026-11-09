package brave

import (
	"github.com/funtimecoding/soil/pkg/brave/profile"
	"github.com/funtimecoding/soil/pkg/errors/not_found"
)

func Profile(name string) (*profile.Profile, error) {
	all, e := Profiles()

	if e != nil {
		return nil, e
	}

	for _, p := range all {
		if p.Profile == name {
			return p, nil
		}
	}

	return nil, not_found.New("profile", name)
}
