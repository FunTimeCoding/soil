package route

import (
	"github.com/funtimecoding/soil/pkg/strings/join"
	"net/http"
)

func Get(path ...string) string {
	return join.Space(http.MethodGet, join.Empty(path...))
}
