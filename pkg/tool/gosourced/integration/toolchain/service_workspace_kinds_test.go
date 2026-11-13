package toolchain

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/lint/analyzer/unclosed_resource"
	"github.com/funtimecoding/soil/pkg/lint/fact"
	"github.com/funtimecoding/soil/pkg/source/index"
	"github.com/funtimecoding/soil/pkg/source/unit/kind_tester"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/service/workspace"
	"testing"
)

func TestWorkspaceHandleCarriesSummaryKind(t *testing.T) {
	directory := kind_tester.Module(t)
	w := workspace.New(t.TempDir()).Workspace(directory)
	assert.Integer(
		t,
		2,
		len(index.Facts[*fact.Summary](w, unclosed_resource.Kind())),
	)
}
