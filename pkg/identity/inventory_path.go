package identity

import (
	"github.com/funtimecoding/soil/pkg/constant"
	"github.com/funtimecoding/soil/pkg/strings/join"
	"path/filepath"
)

func (t *Tool) InventoryPath() string {
	return filepath.Join(
		t.StorageDirectory(false),
		join.Empty(constant.InventoryName, constant.MarkupExtension),
	)
}
