package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/source/index"
	"github.com/funtimecoding/soil/pkg/source/module_graph"
	"github.com/funtimecoding/soil/pkg/source/unit/kind_tester"
	"github.com/funtimecoding/soil/pkg/system/system_tester"
	"testing"
)

func TestIndexWithGraphAnswersOverGivenGraph(t *testing.T) {
	directory := kind_tester.New(t).Module()
	g := module_graph.Build(directory)
	system_tester.New(t).Write(directory, "three/three.go", "package three\n")
	w := index.NewWithGraph(t.TempDir(), directory, g)
	assert.True(t, w.Reaches("example/two"))
	assert.False(t, w.Reaches("example/three"))
	assert.True(t, index.New(t.TempDir(), directory).Reaches("example/three"))
}
