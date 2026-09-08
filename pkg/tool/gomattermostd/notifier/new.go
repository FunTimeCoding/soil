package notifier

import (
	"github.com/funtimecoding/soil/pkg/face"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/connector"
)

func New(
	c *connector.Client,
	source string,
	r face.Reporter,
) *Notifier {
	return &Notifier{connector: c, source: source, reporter: r}
}
