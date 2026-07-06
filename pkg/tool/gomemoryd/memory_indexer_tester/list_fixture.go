package memory_indexer_tester

import (
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/generated/client"
	"github.com/funtimecoding/soil/pkg/web"
	"net/http"
)

func ListFixture(v ...*client.SearchResult) http.HandlerFunc {
	var results []client.SearchResult

	for _, p := range v {
		results = append(results, *p)
	}

	return func(
		w http.ResponseWriter,
		_ *http.Request,
	) {
		web.EncodeNotation(
			w,
			client.ListOutcome{Results: results},
		)
	}
}
