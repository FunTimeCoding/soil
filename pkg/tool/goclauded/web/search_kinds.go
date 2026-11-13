package web

import (
	"github.com/funtimecoding/soil/pkg/tool/goclauded/constant"
	"net/http"
)

func searchKinds(q *http.Request) []string {
	kinds := q.URL.Query()[constant.KindField]

	if len(kinds) == 0 {
		return []string{constant.BlockMessage}
	}

	return kinds
}
