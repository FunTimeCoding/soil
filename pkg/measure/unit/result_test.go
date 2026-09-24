package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/measure/count"
	"github.com/funtimecoding/soil/pkg/measure/file"
	"github.com/funtimecoding/soil/pkg/measure/result"
	"testing"
)

func TestByLanguageSortsByCode(t *testing.T) {
	r := result.New()
	r.Add(file.New("a.go", "Go", counted(2, 1, 0)))
	r.Add(file.New("b.md", "Markdown", counted(9, 0, 1)))
	r.Add(file.New("c.go", "Go", counted(3, 0, 0)))
	v := r.ByLanguage()
	assert.Count(t, 2, v)
	assert.String(t, "Markdown", v[0].Language)
	assert.Integer(t, 1, v[0].Files)
	assert.String(t, "Go", v[1].Language)
	assert.Integer(t, 2, v[1].Files)
	assert.Integer(t, 5, v[1].Count.Code)
	assert.Integer(t, 1, v[1].Count.Comment)
	total := r.Total()
	assert.Integer(t, 3, total.Files)
	assert.Integer(t, 14, total.Count.Code)
	assert.Integer(t, 16, total.Count.Total())
}

func TestByPathSorts(t *testing.T) {
	r := result.New()
	r.Add(file.New("b.go", "Go", count.New()))
	r.Add(file.New("a.go", "Go", count.New()))
	v := r.ByPath()
	assert.String(t, "a.go", v[0].Path)
	assert.String(t, "b.go", v[1].Path)
	assert.String(t, "b.go", r.Files[0].Path)
}
