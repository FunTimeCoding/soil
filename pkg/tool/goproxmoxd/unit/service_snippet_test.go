package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/errors/dropped"
	"github.com/funtimecoding/soil/pkg/errors/not_found"
	"github.com/funtimecoding/soil/pkg/errors/validation"
	"github.com/funtimecoding/soil/pkg/tool/goproxmoxd/constant"
	"github.com/funtimecoding/soil/pkg/tool/goproxmoxd/unit/service_tester"
	"testing"
)

func TestSnippetRoundTripsThroughService(t *testing.T) {
	s := service_tester.New(t).Service()
	volume, e := s.CreateSnippet("alfa", "alfa.yaml", "#cloud-config")
	assert.FatalOnError(t, e)
	assert.String(t, "local:snippets/alfa.yaml", volume)
	b, e := s.GetSnippet("alfa", "alfa.yaml")
	assert.FatalOnError(t, e)
	assert.String(t, "#cloud-config", string(b))
	listed, e := s.ListSnippets("alfa")
	assert.FatalOnError(t, e)
	assert.Count(t, 1, listed)
	assert.String(t, "alfa.yaml", listed[0].Name())
	assert.FatalOnError(t, s.DeleteSnippet("alfa", "alfa.yaml"))
	empty, e := s.ListSnippets("alfa")
	assert.FatalOnError(t, e)
	assert.Count(t, 0, empty)
}

func TestSnippetListingLeavesDirectoriesOut(t *testing.T) {
	o := service_tester.New(t)
	s := o.Service()
	h := o.Alfa().Shell
	h.AddDirectory("/var/lib/vz/snippets/bravo")
	_, e := s.CreateSnippet("alfa", "alfa.yaml", "#cloud-config")
	assert.FatalOnError(t, e)
	held, f := h.ListDirectory(constant.SnippetDirectory)
	assert.FatalOnError(t, f)
	assert.Count(t, 2, held)
	listed, e := s.ListSnippets("alfa")
	assert.FatalOnError(t, e)
	assert.Count(t, 1, listed)
	assert.String(t, "alfa.yaml", listed[0].Name())
}

func TestDroppedShellIsRedialledOnNextCall(t *testing.T) {
	o := service_tester.New(t, service_tester.WithRedial())
	s := o.Service()
	_, e := s.ListSnippets("alfa")
	assert.FatalOnError(t, e)
	first := o.Dialled()[0]
	first.Fail(first.ListDirectory, dropped.Format("connection dropped"))
	_, e = s.ListSnippets("alfa")
	assert.True(t, dropped.Is(e))
	_, e = s.ListSnippets("alfa")
	assert.FatalOnError(t, e)
	assert.Count(t, 2, o.Dialled())
	assert.Integer(t, 1, first.Closed)
}

func TestShellFailingOtherwiseIsKept(t *testing.T) {
	o := service_tester.New(t, service_tester.WithRedial())
	s := o.Service()
	_, e := s.GetSnippet("alfa", "alfa.yaml")
	assert.True(t, not_found.Is(e))
	_, e = s.GetSnippet("alfa", "alfa.yaml")
	assert.True(t, not_found.Is(e))
	assert.Count(t, 1, o.Dialled())
	assert.Integer(t, 0, o.Dialled()[0].Closed)
}

func TestSnippetNameLeavingDirectoryIsRefused(t *testing.T) {
	o := service_tester.New(t)
	s := o.Service()
	h := o.Alfa().Shell
	_, e := s.CreateSnippet("alfa", "../alfa", "#cloud-config")
	assert.True(t, validation.Is(e))
	_, e = s.GetSnippet("alfa", "bravo/../../charlie")
	assert.True(t, validation.Is(e))
	assert.True(t, validation.Is(s.DeleteSnippet("alfa", "/etc/delta")))
	held, f := h.ListDirectory("/var/lib/vz")
	assert.FatalOnError(t, f)
	assert.Count(t, 0, held)
}

func TestMissingSnippetIsNotFound(t *testing.T) {
	s := service_tester.New(t).Service()
	_, e := s.GetSnippet("alfa", "alfa.yaml")
	assert.True(t, not_found.Is(e))
	assert.True(t, not_found.Is(s.DeleteSnippet("alfa", "alfa.yaml")))
}

func TestShellFailureReachesCaller(t *testing.T) {
	o := service_tester.New(t)
	s := o.Service()
	h := o.Alfa().Shell
	h.Fail(h.WriteFile, dropped.Format("connection dropped"))
	_, e := s.CreateSnippet("alfa", "alfa.yaml", "#cloud-config")
	assert.True(t, dropped.Is(e))
	h.Fail(h.ListDirectory, dropped.Format("connection dropped"))
	_, e = s.ListSnippets("alfa")
	assert.True(t, dropped.Is(e))
}

func TestSnippetWithoutNameOrContentIsRefused(t *testing.T) {
	s := service_tester.New(t).Service()
	_, e := s.CreateSnippet("alfa", "", "#cloud-config")
	assert.True(t, validation.Is(e))
	_, e = s.CreateSnippet("alfa", "alfa.yaml", "")
	assert.True(t, validation.Is(e))
}
