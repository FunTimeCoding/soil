package route

import (
	"github.com/funtimecoding/soil/pkg/strings/join"
	"net/http"
)

func Delete(path ...string) string {
	return join.Space(http.MethodDelete, join.Empty(path...))
}
