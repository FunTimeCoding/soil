package tag

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/console/status"
	"github.com/funtimecoding/soil/pkg/console/status/option"
)

func (t *Tag) Format(f *option.Format) string {
	s := status.New(f).Integer(t.Identifier).DetailLink(
		fmt.Sprintf("https://%s/tags/%d", t.Host, t.Identifier),
		t.formatName(f),
		"",
	)

	return s.Format()
}
