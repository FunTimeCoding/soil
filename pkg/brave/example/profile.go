package example

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/brave"
	"github.com/funtimecoding/soil/pkg/brave/constant"
)

func Profile() {
	for _, p := range brave.Profiles() {
		fmt.Printf("Profile: %+v\n", p)

		if false {
			if p.Profile == constant.Profile2 {
				brave.OpenProfile(p.Profile)
			}
		}
	}
}
