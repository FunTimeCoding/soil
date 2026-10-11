package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/service/result/location"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/service/result/references"
	"testing"
)

func TestPaginateReferences(t *testing.T) {
	references := references.New(
		"Used",
		[]*location.Location{
			location.New("a.go", 1, ""),
			location.New("b.go", 2, ""),
			location.New("c.go", 3, ""),
			location.New("d.go", 4, ""),
		},
	)
	references.Paginate(2, 1)
	assert.Integer(t, 4, references.Total)
	assert.Integer(t, 2, len(references.Locations))
	assert.String(t, "b.go", references.Locations[0].File)
	assert.String(t, "c.go", references.Locations[1].File)
	assert.Integer(t, 1, references.More)
}

func TestPaginateReferencesOffsetBeyond(t *testing.T) {
	references := references.New(
		"Used",
		[]*location.Location{location.New("a.go", 1, "")},
	)
	references.Paginate(25, 5)
	assert.Integer(t, 0, len(references.Locations))
	assert.Integer(t, 0, references.More)
}
