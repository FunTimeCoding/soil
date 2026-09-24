package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/tool/gorunif/run_if"
	"github.com/funtimecoding/soil/pkg/tool/gorunif/run_if/option"
	"testing"
)

func TestMatchesPrefix(t *testing.T) {
	o := option.New()
	o.Pattern = "strata/manifest/"
	assert.True(
		t,
		run_if.Matches(o, []string{"pkg/a.go", "strata/manifest/x.yaml"}),
	)
	assert.False(t, run_if.Matches(o, []string{"pkg/strata/manifest/x.yaml"}))
	assert.False(t, run_if.Matches(o, nil))
}

func TestMatchesSuffix(t *testing.T) {
	o := option.New()
	o.Pattern = ".go"
	o.Suffix = true
	assert.True(t, run_if.Matches(o, []string{"doc/x.md", "pkg/a.go"}))
	assert.False(t, run_if.Matches(o, []string{"doc/x.md", "pkg/a.go.txt"}))
}
