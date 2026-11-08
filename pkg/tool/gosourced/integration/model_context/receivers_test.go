package model_context

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/constant"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/integration/model_context_tester"
	"os"
	"path/filepath"
	"testing"
)

func TestRenameReceiversThroughTheTool(t *testing.T) {
	o := model_context_tester.New(
		t,
		"../../service/testdata/rename-receivers/src",
	)
	o.Client.MustCallTool(constant.UseModule, map[string]any{"module": "test"})
	result := o.Client.MustCallTool(
		constant.RenameReceivers,
		map[string]any{
			"package_path": "example/pkg/target",
			"receivers":    []any{map[string]any{"type": "Shelf", "name": "s"}},
		},
	)
	assert.StringContains(t, "Shelf.Name: y → s, 1 uses", result)
	source, e := os.ReadFile(filepath.Join(o.Directory, "pkg/target/shelf.go"))
	assert.FatalOnError(t, e)
	assert.StringContains(t, "return s.product.Name", string(source))
}
