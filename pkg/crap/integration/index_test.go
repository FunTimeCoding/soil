package integration

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/crap/constant"
	"github.com/funtimecoding/soil/pkg/crap/function"
	"github.com/funtimecoding/soil/pkg/crap/index"
	"testing"
)

func TestIndexLoad(t *testing.T) {
	root := module(t)
	i := index.Load(root, constant.AllPackages)
	assert.Count(t, 3, i.Functions)
	half := i.ByKey(function.NewKey("example.test/m/pkg/a", "half.go", 3))
	assert.NotNil(t, half)
	assert.String(t, "Half", half.Name)
	assert.Integer(t, 2, half.Complexity)
	assert.String(t, "example.test/m/pkg/a", half.Package)
	untested := i.ByKey(function.NewKey("example.test/m/pkg/a", "thing.go", 5))
	assert.String(t, "*Thing", untested.Receiver)
	assert.String(t, "*Thing.Untested", untested.QualifiedName())
	assert.Integer(t, 3, untested.Complexity)
	assert.Integer(t, 11, untested.EndLine)
}

func TestKeyMatchesCoverReport(t *testing.T) {
	assert.String(
		t,
		"example.test/m/pkg/a/half.go:3",
		function.NewKey("example.test/m/pkg/a", "/any/where/pkg/a/half.go", 3),
	)
}
