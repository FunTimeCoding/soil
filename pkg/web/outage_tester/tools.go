package outage_tester

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/errors/sentry/reporter/memory"
	"github.com/funtimecoding/soil/pkg/generative/model_context_server"
	"github.com/funtimecoding/soil/pkg/web/constant"
	"slices"
	"strings"
	"testing"
)

func Tools(
	t *testing.T,
	serve func(*memory.Memory) (*model_context_server.Server, func() error),
	usual constant.OutageReach,
	exceptions map[string]constant.OutageReach,
	overrides map[string]map[string]any,
) {
	t.Helper()
	r := memory.New()
	listed, _ := serve(r)
	listing := listed.Connect(t)
	tools := listing.ListTools()
	listing.Close()
	listed.Close()

	if len(tools) == 0 {
		t.Fatal("the server lists no tools")
	}

	var served []string
	var problems []string

	for _, tool := range tools {
		served = append(served, tool.Name)
		want, okay := exceptions[tool.Name]

		if !okay {
			want = usual
		}

		before := len(r.Events())
		s, unavailable := serve(r)
		c := s.Connect(t)
		answer, e := c.CallToolRaw(
			tool.Name,
			arguments(tool, overrides[tool.Name]),
		)
		c.Close()
		outage := unavailable()
		s.Close()

		if e != nil {
			problems = append(
				problems,
				fmt.Sprintf("%s: no answer: %v", tool.Name, e),
			)

			continue
		}

		said := text(answer)
		problems = append(
			problems,
			verdict(
				tool.Name,
				want,
				answer.IsError,
				outage != nil && strings.Contains(said, outage.Error()),
				outage != nil,
				len(r.Events())-before,
				said,
			)...,
		)
	}

	problems = append(problems, declared(served, exceptions, overrides)...)
	slices.Sort(problems)
	assert.String(t, "", strings.Join(problems, ", "))
}
