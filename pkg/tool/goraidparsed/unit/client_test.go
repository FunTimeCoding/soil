package unit

import (
	"errors"
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/errors/not_found"
	"github.com/funtimecoding/soil/pkg/errors/unexpected"
	"github.com/funtimecoding/soil/pkg/tool/goraidparsed/types/report"
	"github.com/funtimecoding/soil/pkg/tool/goraidparsed/unit/fixed_generator"
	"testing"
)

func TestGenerateReturnsTheReport(t *testing.T) {
	c := mountedClient(t, fixed_generator.New(report.New("bravo.html", 2), nil))
	result, e := c.PostGenerate(t.Context(), []string{"charlie.json"}, nil)
	assert.FatalOnError(t, e)
	assert.String(t, "bravo.html", result.Path)
	assert.Integer(t, 2, result.TidCount)
	assert.FatalOnError(t, c.Reachable(t.Context()))
}

func TestAMissingInputFileIsNotFound(t *testing.T) {
	c := mountedClient(
		t,
		fixed_generator.New(nil, not_found.New("input file", "delta.json")),
	)
	_, e := c.PostGenerate(t.Context(), []string{"delta.json"}, nil)
	assert.True(t, not_found.Is(e))
	assert.StringContains(t, "delta.json", e.Error())
}

func TestAFailedGenerateIsUnexpected(t *testing.T) {
	c := mountedClient(
		t,
		fixed_generator.New(nil, errors.New("echo parse failed")),
	)
	_, e := c.PostGenerate(t.Context(), []string{"charlie.json"}, nil)
	assert.True(t, unexpected.Is(e))
}
