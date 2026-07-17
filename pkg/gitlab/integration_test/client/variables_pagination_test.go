//go:build local

package client

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/gitlab/integration_test/client_tester"
	"github.com/funtimecoding/soil/pkg/web"
	"net/http"
	"testing"
)

func TestVariablesPagination(t *testing.T) {
	var pages []string
	r := client_tester.New(
		t,
		func(m *http.ServeMux) {
			m.HandleFunc(
				"/api/v4/projects/7/variables",
				func(w http.ResponseWriter, q *http.Request) {
					page := q.URL.Query().Get("page")
					pages = append(pages, page)
					start, count := 0, 100

					if page == "2" {
						start, count = 100, 50
					}

					var list []map[string]string

					for i := start; i < start+count; i++ {
						list = append(
							list,
							map[string]string{
								"key":   fmt.Sprintf("KEY_%d", i),
								"value": "v",
							},
						)
					}

					web.Encode(w, list)
				},
			)
		},
	)
	result, e := r.Client.Variables(7)

	if e != nil {
		t.Fatal(e)
	}

	assert.Integer(t, 150, len(result))
	seen := make(map[string]bool)

	for _, v := range result {
		seen[v.Key] = true
	}

	assert.Integer(t, 150, len(seen))
	assert.Integer(t, 2, len(pages))
	assert.String(t, "1", pages[0])
	assert.String(t, "2", pages[1])
}
