package route

import (
	"github.com/funtimecoding/soil/pkg/strings/join"
	"net/http"
)

func Post(path ...string) string {
	return join.Space(http.MethodPost, join.Empty(path...))
}
