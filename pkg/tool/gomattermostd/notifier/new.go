package notifier

import (
	"github.com/funtimecoding/soil/pkg/face"
	goclauded "github.com/funtimecoding/soil/pkg/tool/goclauded/face"
)

func New(
	c goclauded.Client,
	source string,
	r face.Reporter,
) *Notifier {
	return &Notifier{connector: c, source: source, reporter: r}
}
