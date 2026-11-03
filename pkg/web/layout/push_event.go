package layout

import (
	"bytes"
	"github.com/funtimecoding/soil/pkg/errors"
	stringsConstant "github.com/funtimecoding/soil/pkg/strings/constant"
	"github.com/funtimecoding/soil/pkg/strings/split"
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
	emit(w, "event: %s\n", name)

	for _, l := range split.NewLine(b.String()) {
		emit(w, "data: %s\n", l)
	}

	emit(w, stringsConstant.Unix)
}
