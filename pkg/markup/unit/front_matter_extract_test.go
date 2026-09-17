package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/markup/front_matter"
	"github.com/funtimecoding/soil/pkg/markup/scalar_or_list"
	"testing"
)

func TestFrontMatterExtract(t *testing.T) {
	f, found := front_matter.Extract(
		"---\ntitle: Alfa\nbase: pkg/lint\n---\n# Alfa\n",
	)
	assert.True(t, found)
	assert.String(t, "title: Alfa\nbase: pkg/lint", f.Raw)
	assert.Integer(t, 4, f.Lines)
	assert.Integer(t, 3, f.Line("base"))
	assert.Integer(t, 0, f.Line("world"))
}

func TestFrontMatterExtractIgnoresLongerRule(t *testing.T) {
	f, found := front_matter.Extract("---\ntitle: Alfa\n----\nbody\n---\n")
	assert.True(t, found)
	assert.String(t, "title: Alfa\n----\nbody", f.Raw)
	assert.Integer(t, 5, f.Lines)
}

func TestFrontMatterExtractAbsent(t *testing.T) {
	_, found := front_matter.Extract("# Alfa\n")
	assert.Boolean(t, false, found)
	_, found = front_matter.Extract("---\ntitle: Alfa\n")
	assert.Boolean(t, false, found)
}

func TestFrontMatterDecodeScalarOrList(t *testing.T) {
	f, _ := front_matter.Extract("---\nbase: pkg/lint, pkg/markup\n---\n")
	var scalar declaration
	assert.Nil(t, f.Decode(&scalar))
	assert.Strings(t, []string{"pkg/lint, pkg/markup"}, scalar.Base)
	f, _ = front_matter.Extract(
		"---\nbase:\n  - pkg/lint\n  - pkg/markup\n---\n",
	)
	var list declaration
	assert.Nil(t, f.Decode(&list))
	assert.Strings(t, []string{"pkg/lint", "pkg/markup"}, list.Base)
}

type declaration struct {
	Base scalar_or_list.Strings `yaml:"base"`
}
