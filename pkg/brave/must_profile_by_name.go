package brave

import (
	"github.com/funtimecoding/soil/pkg/brave/profile"
	"github.com/funtimecoding/soil/pkg/errors"
)

func MustProfileByName(name string) *profile.Profile {
	p := ProfileByName(name)

	if p == nil {
		errors.NotFound(name)

		return profile.New()
	}

	return p
}
