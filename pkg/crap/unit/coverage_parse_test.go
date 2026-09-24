package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/crap/coverage"
	"testing"
)

func TestParseFunctionReport(t *testing.T) {
	v := coverage.Parse(
		`github.com/x/m/pkg/a/add.go:3:		Add		100.0%
github.com/x/m/pkg/a/run.go:10:		Run		0.0%
github.com/x/m/pkg/a/walk.go:11:		walk		95.2%
total:			(statements)		87.9%
`,
	)
	assert.Count(t, 3, v)
	assert.Float(t, 100, v["github.com/x/m/pkg/a/add.go:3"])
	assert.Float(t, 0, v["github.com/x/m/pkg/a/run.go:10"])
	assert.Float(t, 95.2, v["github.com/x/m/pkg/a/walk.go:11"])
}

func TestParseEmpty(t *testing.T) {
	assert.Count(t, 0, coverage.Parse(""))
}
