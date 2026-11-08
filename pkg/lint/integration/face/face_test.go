package face

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/lint/face"
	"github.com/funtimecoding/soil/pkg/lint/unit/face_tester"
	"testing"
)

func TestUniverseErrorIsImplemented(t *testing.T) {
	o := face_tester.New(t)
	assert.True(t, face.New(o.Loaded).Implements(o.Method("Error")))
}

func TestDirectlyImportedInterfaceIsImplemented(t *testing.T) {
	o := face_tester.New(t)
	assert.True(t, face.New(o.Loaded).Implements(o.Method("MarshalJSON")))
}

func TestIndexSeesAnInterfaceBehindItsImporter(t *testing.T) {
	o := face_tester.New(t)
	assert.True(
		t,
		face.NewWithIndex(o.Workspace, o.Loaded).Implements(
			o.Method("MarshalText"),
		),
	)
}

func TestWorkspaceSeesAnInterfaceBehindItsImporter(t *testing.T) {
	o := face_tester.New(t)
	assert.True(
		t,
		face.FromWorkspace(o.Workspace, o.Loaded).Implements(
			o.Method("MarshalText"),
		),
	)
}

func TestNameWithAnotherSignatureImplementsNothing(t *testing.T) {
	o := face_tester.New(t)
	assert.False(
		t,
		face.NewWithIndex(o.Workspace, o.Loaded).Implements(o.Method("Write")),
	)
}
