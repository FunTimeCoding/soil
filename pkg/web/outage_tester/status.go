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
	"strings"
)

func (o *Tester) status(
	s *model_context_server.Server,
	path string,
	method string,
	operation *openapi3.Operation,
	parameters map[string]string,
	bodies map[string]string,
) (int, error) {
	o.t.Helper()
	var content io.Reader

	if method != http.MethodGet && method != http.MethodDelete {
		v, okay := bodies[operation.OperationID]

		if !okay {
			v = body(operation)
		}

		content = strings.NewReader(v)
	}

	q, e := http.NewRequestWithContext(
		o.t.Context(),
		method,
		s.Address(
			join.Empty(
				fillPath(path, parameters),
				query(operation, parameters),
			),
		),
		content,
	)
	assert.FatalOnError(o.t, e)
	web.Bearer(q, generativeConstant.ModelContextTestToken)
	q.Header.Set(constant.ContentType, constant.Object)
	r, f := s.Client().Do(q)

	if f != nil {
		return 0, f
	}

	errors.PanicClose(r.Body)

	return r.StatusCode, nil
}
