package example

import (
	"github.com/funtimecoding/soil/pkg/brave"
	"github.com/funtimecoding/soil/pkg/brave/constant"
	"github.com/funtimecoding/soil/pkg/console"
)

func Profile() {
	for _, p := range brave.Profiles() {
		console.Format("Profile: %+v\n", p)

		if false {
			if p.Profile == constant.Profile2 {
				brave.OpenProfile(p.Profile)
			}
		}
	}
}
