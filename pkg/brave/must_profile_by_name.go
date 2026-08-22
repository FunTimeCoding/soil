package brave

import (
	"github.com/funtimecoding/soil/pkg/brave/profile"
	"github.com/funtimecoding/soil/pkg/errors/not_found"
)

func MustProfileByName(name string) *profile.Profile {
	p := ProfileByName(name)

	if p == nil {
		panic(not_found.New("profile", name))
	}

	return p
}
