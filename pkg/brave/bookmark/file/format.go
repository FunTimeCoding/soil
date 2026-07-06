package file

import (
	"github.com/funtimecoding/soil/pkg/console/status"
	"github.com/funtimecoding/soil/pkg/console/status/option"
)

func (n *Node) Format(f *option.Format) string {
	s := status.New(f).String(
		n.Type,
		n.formatName(f),
		n.Locator,
	).RawList(n)

	return s.Format()
}
