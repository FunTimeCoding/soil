package coverage

import (
	"github.com/funtimecoding/soil/pkg/generative/anthropic/claude/tool_call"
	generative "github.com/funtimecoding/soil/pkg/generative/constant"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/constant"
	"sort"
	"strings"
	"time"
)

func Compute(
	calls []tool_call.Call,
	registered map[string][]string,
	paths map[string]string,
	configured map[string]bool,
	aliases map[string]string,
	now time.Time,
) []*Server {
	cut := now.Add(-constant.CoverageRecentWindow)
	tools := map[string]map[string]*Tool{}
	ensure := func(
		server string,
		name string,
	) *Tool {
		if tools[server] == nil {
			tools[server] = map[string]*Tool{}
		}

		if tools[server][name] == nil {
			tools[server][name] = &Tool{Name: name}
		}

		return tools[server][name]
	}

	for server, names := range registered {
		if tools[server] == nil {
			tools[server] = map[string]*Tool{}
		}

		for _, name := range names {
			ensure(server, name).Registered = true
		}
	}

	for server := range configured {
		if tools[server] == nil {
			tools[server] = map[string]*Tool{}
		}
	}

	for _, c := range calls {
		if !strings.HasPrefix(c.Name, generative.ModelContextToolPrefix) {
			continue
		}

		server, name, found := strings.Cut(
			strings.TrimPrefix(c.Name, generative.ModelContextToolPrefix),
			generative.ModelContextToolSeparator,
		)

		if !found {
			continue
		}

		if canonical, aliased := aliases[server]; aliased {
			server = canonical
		}

		t := ensure(server, name)
		t.CallsTotal++
		when, e := time.Parse(time.RFC3339Nano, c.Timestamp)

		if e != nil {
			continue
		}

		if !when.Before(cut) {
			t.CallsRecent++
		}

		if when.After(t.LastUsed) {
			t.LastUsed = when
		}
	}

	var result []*Server

	for name, byTool := range tools {
		s := &Server{
			Name:       name,
			Path:       paths[name],
			Configured: configured[name],
			Registered: len(registered[name]),
		}

		for _, t := range byTool {
			s.Tools = append(s.Tools, t)
			s.CallsTotal += t.CallsTotal
			s.CallsRecent += t.CallsRecent

			if t.Registered && t.CallsTotal > 0 {
				s.UsedTotal++
			}

			if t.Registered && t.CallsRecent > 0 {
				s.UsedRecent++
			}

			if t.LastUsed.After(s.LastUsed) {
				s.LastUsed = t.LastUsed
			}
		}

		sort.Slice(
			s.Tools,
			func(i, j int) bool {
				if s.Tools[i].CallsTotal != s.Tools[j].CallsTotal {
					return s.Tools[i].CallsTotal > s.Tools[j].CallsTotal
				}

				return s.Tools[i].Name < s.Tools[j].Name
			},
		)
		result = append(result, s)
	}

	sort.Slice(
		result,
		func(i, j int) bool {
			return result[i].Name < result[j].Name
		},
	)

	return result
}
