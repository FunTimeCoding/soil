package outage_tester

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/errors"
	generativeConstant "github.com/funtimecoding/soil/pkg/generative/constant"
	"github.com/funtimecoding/soil/pkg/generative/model_context_server"
	"github.com/funtimecoding/soil/pkg/strings/join"
	"github.com/funtimecoding/soil/pkg/web"
	"github.com/funtimecoding/soil/pkg/web/constant"
	"github.com/getkin/kin-openapi/openapi3"
	"io"
	"net/http"
	"regexp"
	"strings"
	"testing"
)

func Statuses(
	t *testing.T,
	serve func() *model_context_server.Server,
	spec *openapi3.T,
	bodies map[string]string,
) map[string]int {
	t.Helper()
	parameter := regexp.MustCompile(`\{[^}]+\}`)
	result := map[string]int{}

	for path, item := range spec.Paths.Map() {
		for method, o := range item.Operations() {
			var content io.Reader

			if method != http.MethodGet && method != http.MethodDelete {
				v, okay := bodies[o.OperationID]

				if !okay {
					v = body(o)
				}

				content = strings.NewReader(v)
			}

			s := serve()
			q, e := http.NewRequestWithContext(
				t.Context(),
				method,
				s.Address(
					join.Empty(parameter.ReplaceAllString(path, "1"), query(o)),
				),
				content,
			)
			assert.FatalOnError(t, e)
			web.Bearer(q, generativeConstant.ModelContextTestToken)
			q.Header.Set(constant.ContentType, constant.Object)
			r, f := s.Client().Do(q)
			assert.FatalOnError(t, f)
			errors.PanicClose(r.Body)
			s.Close()
			result[o.OperationID] = r.StatusCode
		}
	}

	if len(result) == 0 {
		t.Fatal("the specification lists no operations")
	}

	return result
}
