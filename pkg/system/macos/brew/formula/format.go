package formula

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/console/status"
	"github.com/funtimecoding/soil/pkg/console/status/option"
	"github.com/funtimecoding/soil/pkg/strings/join"
)

func (f *Formula) Format(o *option.Format) string {
	s := status.New(o).String(
		f.Name,
		fmt.Sprintf(
			"%s → %s",
			join.CommaSpace(f.InstalledVersions),
			f.CurrentVersion,
		),
	)
	s.DetailLink(f.Link, "Homebrew", "")

	return s.Format()
}
