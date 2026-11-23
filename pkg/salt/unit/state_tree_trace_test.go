package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/assert/fixture"
	"github.com/funtimecoding/soil/pkg/errors/not_found"
	"github.com/funtimecoding/soil/pkg/salt/constant"
	"github.com/funtimecoding/soil/pkg/salt/state_tree"
	"github.com/funtimecoding/soil/pkg/salt/types/top_entry"
	"github.com/funtimecoding/soil/pkg/salt/types/trace_report"
	systemConstant "github.com/funtimecoding/soil/pkg/system/constant"
	"testing"
)

func TestTraceStates(t *testing.T) {
	r, e := state_tree.New(
		[]string{fixture.Path(systemConstant.SaltPath, "state")},
		constant.DefaultEnvironments,
	).Trace("web*")
	assert.FatalOnError(t, e)
	assert.Any(
		t,
		[]*trace_report.Node{
			{Name: "helpers.gamma", Orphaned: true},
			{Name: "required", Orphaned: false},
			{Name: "service.alpha", Orphaned: true},
			{Name: "service.beta", Orphaned: true},
			{Name: "web", Orphaned: true},
		},
		r.States,
	)
	assert.Strings(t, []string{"absent"}, r.Missing)
}

func TestTraceFiles(t *testing.T) {
	r, e := state_tree.New(
		[]string{fixture.Path(systemConstant.SaltPath, "state")},
		constant.DefaultEnvironments,
	).Trace("web*")
	assert.FatalOnError(t, e)
	assert.Any(
		t,
		[]*trace_report.Node{
			{Name: "macros/helper.jinja", Orphaned: true},
			{Name: "required/limits.conf", Orphaned: false},
			{Name: "web/conf/site.conf", Orphaned: true},
		},
		r.Files,
	)
}

func TestTraceDeadTarget(t *testing.T) {
	r, e := state_tree.New(
		[]string{fixture.Path(systemConstant.SaltPath, "state")},
		constant.DefaultEnvironments,
	).Trace("st*")
	assert.FatalOnError(t, e)
	assert.Any(
		t,
		[]*top_entry.Entry{
			{
				Environment: "staging",
				Target:      "st*",
				State:       "stagingonly",
				Live:        false,
			},
		},
		r.Entries,
	)
	assert.Any(
		t,
		[]*trace_report.Node{
			{Name: "conditional", Orphaned: false},
			{Name: "stagingonly", Orphaned: true},
		},
		r.States,
	)
}

func TestTraceIgnoresDeadEntries(t *testing.T) {
	r, e := state_tree.New(
		[]string{fixture.Path(systemConstant.SaltPath, "state")},
		constant.DefaultEnvironments,
	).Trace("db*")
	assert.FatalOnError(t, e)
	assert.Any(
		t,
		[]*trace_report.Node{{Name: "conditional", Orphaned: true}},
		r.States,
	)
}

func TestTraceUnknownTarget(t *testing.T) {
	_, e := state_tree.New(
		[]string{fixture.Path(systemConstant.SaltPath, "state")},
		constant.DefaultEnvironments,
	).Trace("nothing*")
	assert.True(t, not_found.Is(e))
}
