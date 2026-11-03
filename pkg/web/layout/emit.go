package layout

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/web"
	"net/http"
)

func emit(
	w http.ResponseWriter,
	format string,
	a ...any,
) {
	_, e := fmt.Fprintf(w, format, a...)
	web.PanicOnWriteError(e)
}
