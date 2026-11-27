package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/browser/ladybird/types/request"
	"github.com/funtimecoding/soil/pkg/errors/not_found"
	"github.com/funtimecoding/soil/pkg/errors/validation"
	"github.com/funtimecoding/soil/pkg/strings/join"
	"github.com/funtimecoding/soil/pkg/tool/goladybirdd/constant"
	"github.com/funtimecoding/soil/pkg/tool/goladybirdd/types/fill"
	"github.com/funtimecoding/soil/pkg/tool/goladybirdd/unit/service_tester"
	"os"
	"strings"
	"testing"
)

func TestSnapshotFormatsTree(t *testing.T) {
	o := service_tester.New(t)
	o.Client.AddContext("alfa", "about:blank")
	o.Page("Probe", "Press", "Name")
	v, e := o.Service().Snapshot("")
	assert.FatalOnError(t, e)
	assert.String(
		t,
		"[document] Probe\n  uid=e1 [button] Press\n  uid=e2 [button] Name\n  [text] idle\n",
		v.Text,
	)
	assert.String(t, "", v.Path)
}

func TestLargeSnapshotSavedToFile(t *testing.T) {
	o := service_tester.New(t)
	o.Client.AddContext("alfa", "about:blank")
	o.Page(
		"Probe",
		strings.Split(
			strings.Repeat(join.Empty(strings.Repeat("x", 60), ","), 400),
			",",
		)...,
	)
	v, e := o.Service().Snapshot("")
	assert.FatalOnError(t, e)
	assert.String(t, "", v.Text)
	assert.Suffix(t, "ladybird_alfa_snapshot.txt", v.Path)
	assert.Integer(t, 403, v.Lines)
	b, f := os.ReadFile(v.Path)
	assert.FatalOnError(t, f)
	assert.Prefix(t, "[document] Probe\n  uid=e1 [button] xxx", string(b))
}

func TestClickUsesSnapshotReference(t *testing.T) {
	o := service_tester.New(t)
	o.Client.AddContext("alfa", "about:blank")
	o.Page("Probe", "Press", "Name")
	_, e := o.Service().Snapshot("")
	assert.FatalOnError(t, e)
	_, e = o.Service().Click("", "e2", false)
	assert.FatalOnError(t, e)
	calls := o.Client.Calls()
	last := calls[len(calls)-1]
	assert.String(t, "(e) => e.click()", last.Declaration)
	assert.String(
		t,
		"ref-1",
		last.Arguments[0].(*request.Reference).SharedIdentifier,
	)
}

func TestClickBeforeSnapshotNotFound(t *testing.T) {
	o := service_tester.New(t)
	o.Client.AddContext("alfa", "about:blank")
	_, e := o.Service().Click("", "e1", false)
	assert.True(t, not_found.Is(e))
	assert.String(t, "uid e1 not found - take a snapshot first", e.Error())
}

func TestClickGoneElementNotFound(t *testing.T) {
	o := service_tester.New(t)
	o.Client.AddContext("alfa", "about:blank")
	o.Page("Probe", "Press")
	_, e := o.Service().Snapshot("")
	assert.FatalOnError(t, e)
	o.Client.Refuse(
		constant.ClickScript,
		not_found.Format("no such node: Node not found"),
	)
	_, e = o.Service().Click("", "e1", false)
	assert.True(t, not_found.Is(e))
	assert.String(
		t,
		"uid e1 is gone from the page - take a new snapshot",
		e.Error(),
	)
}

func TestClickTakesNewSnapshot(t *testing.T) {
	o := service_tester.New(t)
	o.Client.AddContext("alfa", "about:blank")
	o.Page("Probe", "Press")
	_, e := o.Service().Snapshot("")
	assert.FatalOnError(t, e)
	v, f := o.Service().Click("", "e1", true)
	assert.FatalOnError(t, f)
	assert.String(
		t,
		"[document] Probe\n  uid=e1 [button] Press\n  [text] idle\n",
		v.Text,
	)
}

func TestFillTypesValue(t *testing.T) {
	o := service_tester.New(t)
	o.Client.AddContext("alfa", "about:blank")
	o.Page("Probe", "Press", "Name")
	_, e := o.Service().Snapshot("")
	assert.FatalOnError(t, e)
	r := fill.New("", "e2", "Charlie")
	r.Snapshot = false
	_, e = o.Service().Fill(r)
	assert.FatalOnError(t, e)
	calls := o.Client.Calls()
	last := calls[len(calls)-1]
	assert.String(
		t,
		"(e) => { e.focus(); if (e.select) e.select(); }",
		last.Declaration,
	)
	assert.Strings(t, []string{"Charlie"}, o.Client.Typed())
}

func TestFillDirectSetsValue(t *testing.T) {
	o := service_tester.New(t)
	o.Client.AddContext("alfa", "about:blank")
	o.Page("Probe", "Press", "Name")
	_, e := o.Service().Snapshot("")
	assert.FatalOnError(t, e)
	r := fill.New("", "e2", "Charlie")
	r.Snapshot = false
	r.Direct = true
	_, e = o.Service().Fill(r)
	assert.FatalOnError(t, e)
	calls := o.Client.Calls()
	last := calls[len(calls)-1]
	assert.String(
		t,
		"ref-1",
		last.Arguments[0].(*request.Reference).SharedIdentifier,
	)
	assert.String(t, "Charlie", last.Arguments[1].(*request.Text).Value)
	assert.Count(t, 0, o.Client.Typed())
}

func TestFillRequiresValue(t *testing.T) {
	o := service_tester.New(t)
	o.Client.AddContext("alfa", "about:blank")
	_, e := o.Service().Fill(fill.New("", "e1", ""))
	assert.True(t, validation.Is(e))
}
