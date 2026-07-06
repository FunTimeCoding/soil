package user

import (
	"github.com/funtimecoding/soil/pkg/console/status"
	"github.com/funtimecoding/soil/pkg/console/status/option"
)

func (u *User) Format(f *option.Format) string {
	r := status.New(f).String(u.Name, u.Identifier).RawList(u.Raw)

	if u.FullName != "" {
		r.String(u.FullName)
	}

	return r.Format()
}
