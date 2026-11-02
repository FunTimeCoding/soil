package example

import (
	"github.com/funtimecoding/soil/pkg/brave"
	"github.com/funtimecoding/soil/pkg/brave/constant"
	"github.com/funtimecoding/soil/pkg/console"
	"github.com/funtimecoding/soil/pkg/errors"
)

func Profile() {
	profiles, e := brave.Profiles()
	errors.PanicOnError(e)

	for _, p := range profiles {
		console.Format("Profile: %+v\n", p)

		if false {
			if p.Profile == constant.Profile2 {
				brave.OpenProfile(p.Profile)
			}
		}
	}
}
