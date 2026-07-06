package user

import (
	"github.com/funtimecoding/soil/pkg/console/status"
	"github.com/funtimecoding/soil/pkg/console/status/option"
)

func (u *User) Format(f *option.Format) string {
	return status.New(f).Integer64(u.Identifier).String(u.Name).Format()
}
