package layout

import (
	"bytes"
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/strings/separator"
	"github.com/funtimecoding/soil/pkg/strings/split"
	"github.com/funtimecoding/soil/pkg/system/writer"
	"maragu.dev/gomponents"
	"net/http"
)

func PushEvent(
	w http.ResponseWriter,
	name string,
	n gomponents.Node,
) {
	var b bytes.Buffer
	errors.PanicOnError(n.Render(&b))
	writer.Print(w, "event: %s\n", name)

	for _, l := range split.NewLine(b.String()) {
		writer.Print(w, "data: %s\n", l)
	}

	writer.Print(w, separator.Unix)
}
