package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/errors/not_selected"
	"github.com/funtimecoding/soil/pkg/errors/sentry/reporter/memory"
	"github.com/funtimecoding/soil/pkg/generative/model_context/selection"
	"github.com/funtimecoding/soil/pkg/generative/unit/session_tester"
	"github.com/funtimecoding/soil/pkg/inventory/instance_tool"
	"github.com/funtimecoding/soil/pkg/inventory/instance_tool/argument"
	"github.com/funtimecoding/soil/pkg/inventory/instance_tool/instance_response"
	"github.com/funtimecoding/soil/pkg/inventory/unit/instance_tester"
	"github.com/funtimecoding/soil/pkg/notation"
	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"
	"slices"
	"testing"
)

func TestInstanceToolStaysHiddenWithOneInstance(t *testing.T) {
	s := server.NewMCPServer("test", "0")
	instance_tool.New(
		instance_tester.Inventory("alfa"),
		selection.New(),
		memory.New(),
	).Register(s)
	assert.Count(t, 0, s.ListTools())
}

func TestInstanceToolRegistersWithTwoInstances(t *testing.T) {
	s := server.NewMCPServer("test", "0")
	instance_tool.New(
		instance_tester.Inventory("alfa", "bravo"),
		selection.New(),
		memory.New(),
	).Register(s)
	var names []string

	for name := range s.ListTools() {
		names = append(names, name)
	}

	slices.Sort(names)
	assert.Strings(t, []string{"list_instances", "use_instance"}, names)
}

func TestInstanceToolUseInstanceSelectsForTheSession(t *testing.T) {
	s := server.NewMCPServer("test", "0")
	v := instance_tool.New(
		instance_tester.Inventory("alfa", "bravo"),
		selection.New(),
		memory.New(),
	)
	one := s.WithContext(t.Context(), session_tester.New("one"))
	two := s.WithContext(t.Context(), session_tester.New("two"))
	r, e := v.UseInstance(
		one,
		mcp.CallToolRequest{},
		argument.UseInstance{Instance: "bravo"},
	)
	assert.FatalOnError(t, e)
	assert.False(t, r.IsError)
	selected, f := v.Resolve(one)
	assert.FatalOnError(t, f)
	assert.String(t, "bravo", selected)
	_, g := v.Resolve(two)
	assert.True(t, not_selected.Is(g))
}

func TestInstanceToolUseInstanceRefusesEmptyAndUnknown(t *testing.T) {
	s := server.NewMCPServer("test", "0")
	v := instance_tool.New(
		instance_tester.Inventory("alfa", "bravo"),
		selection.New(),
		memory.New(),
	)
	x := s.WithContext(t.Context(), session_tester.New("one"))
	empty, e := v.UseInstance(x, mcp.CallToolRequest{}, argument.UseInstance{})
	assert.FatalOnError(t, e)
	assert.String(
		t,
		"instance is required",
		empty.Content[0].(mcp.TextContent).Text,
	)
	unknown, f := v.UseInstance(
		x,
		mcp.CallToolRequest{},
		argument.UseInstance{Instance: "charlie"},
	)
	assert.FatalOnError(t, f)
	assert.String(
		t,
		"unknown instance: charlie",
		unknown.Content[0].(mcp.TextContent).Text,
	)
}

func TestInstanceToolCapturesACallWithoutASession(t *testing.T) {
	m := memory.New()
	v := instance_tool.New(
		instance_tester.Inventory("alfa", "bravo"),
		selection.New(),
		m,
	)
	r, e := v.UseInstance(
		t.Context(),
		mcp.CallToolRequest{},
		argument.UseInstance{Instance: "bravo"},
	)
	assert.FatalOnError(t, e)
	assert.True(t, r.IsError)
	assert.Count(t, 1, m.Events())
	_, f := v.Resolve(t.Context())
	assert.False(t, not_selected.Is(f))
	assert.StringContains(t, "no model context session", f.Error())
}

func TestInstanceToolListMarksTheActiveInstance(t *testing.T) {
	s := server.NewMCPServer("test", "0")
	v := instance_tool.New(
		instance_tester.Inventory("alfa", "bravo"),
		selection.New(),
		memory.New(),
	)
	x := s.WithContext(t.Context(), session_tester.New("one"))
	_, e := v.UseInstance(
		x,
		mcp.CallToolRequest{},
		argument.UseInstance{Instance: "bravo"},
	)
	assert.FatalOnError(t, e)
	r, f := v.ListInstances(x, mcp.CallToolRequest{}, argument.ListInstances{})
	assert.FatalOnError(t, f)
	var listed []instance_response.Response
	notation.MustDecodeBytes(
		[]byte(r.Content[0].(mcp.TextContent).Text),
		&listed,
		false,
	)
	assert.Count(t, 2, listed)
	assert.String(t, "alfa", listed[0].Name)
	assert.String(t, "alfa.example", listed[0].Host)
	assert.False(t, listed[0].Active)
	assert.True(t, listed[1].Active)
}
