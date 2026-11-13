package outage_tester

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/errors/sentry/reporter/memory"
	"github.com/funtimecoding/soil/pkg/generative/model_context_server"
	"github.com/funtimecoding/soil/pkg/web/constant"
	"github.com/getkin/kin-openapi/openapi3"
	"net/http"
	"slices"
	"strings"
	"testing"
)

func Operations(
	t *testing.T,
	serve func(*memory.Memory) (*model_context_server.Server, func() error),
	spec *openapi3.T,
	usual constant.OutageReach,
	exceptions map[string]constant.OutageReach,
	parameters map[string]string,
	bodies map[string]string,
) {
	t.Helper()
	r := memory.New()
	var served []string
	var problems []string

	for path, item := range spec.Paths.Map() {
		for method, o := range item.Operations() {
			served = append(served, o.OperationID)
			want, okay := exceptions[o.OperationID]

			if !okay {
				want = usual
			}

			before := len(r.Events())
			s, unavailable := serve(r)
			answered, e := status(t, s, path, method, o, parameters, bodies)
			reached := unavailable() != nil
			s.Close()

			if e != nil {
				problems = append(
					problems,
					fmt.Sprintf("%s: no answer: %v", o.OperationID, e),
				)

				continue
			}

			problems = append(
				problems,
				verdict(
					o.OperationID,
					want,
					answered == http.StatusServiceUnavailable,
					answered == http.StatusServiceUnavailable,
					reached,
					len(r.Events())-before,
					http.StatusText(answered),
				)...,
			)
		}
	}

	if len(served) == 0 {
		t.Fatal("the specification lists no operations")
	}

	problems = append(problems, declared(served, exceptions, bodies)...)
	slices.Sort(problems)
	assert.String(t, "", strings.Join(problems, ", "))
}
