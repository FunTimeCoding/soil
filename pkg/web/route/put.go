package route

import (
	"github.com/funtimecoding/soil/pkg/strings/join"
	"net/http"
)

func Put(path ...string) string {
	return join.Space(http.MethodPut, join.Empty(path...))
}
