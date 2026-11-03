package layout

import (
	stringsConstant "github.com/funtimecoding/soil/pkg/strings/constant"
	"github.com/funtimecoding/soil/pkg/strings/split"
	"net/http"
)

func PushPayload(
	w http.ResponseWriter,
	name string,
	identifier string,
	payload string,
) {
	if identifier != "" {
		emit(w, "id: %s\n", identifier)
	}

	emit(w, "event: %s\n", name)

	for _, l := range split.NewLine(payload) {
		emit(w, "data: %s\n", l)
	}

	emit(w, stringsConstant.Unix)
}
