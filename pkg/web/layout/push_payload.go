package layout

import (
	"github.com/funtimecoding/soil/pkg/strings/constant"
	"github.com/funtimecoding/soil/pkg/strings/split"
	"github.com/funtimecoding/soil/pkg/system/writer"
	"net/http"
)

func PushPayload(
	w http.ResponseWriter,
	name string,
	identifier string,
	payload string,
) {
	if identifier != "" {
		writer.Print(w, "id: %s\n", identifier)
	}

	writer.Print(w, "event: %s\n", name)

	for _, l := range split.NewLine(payload) {
		writer.Print(w, "data: %s\n", l)
	}

	writer.Print(w, constant.Unix)
}
