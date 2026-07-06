package user

import (
	"github.com/funtimecoding/soil/pkg/console"
	"github.com/funtimecoding/soil/pkg/console/status/option"
)

func (u *User) formatName(f *option.Format) string {
	if f.UseColor {
		return console.Cyan("%s", u.Name)
	}

	return u.Name
}
