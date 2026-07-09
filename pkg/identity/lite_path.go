package identity

import (
	"github.com/funtimecoding/soil/pkg/constant"
	"github.com/funtimecoding/soil/pkg/strings/join"
	"path/filepath"
)

func (t *Tool) LitePath() string {
	return filepath.Join(
		t.StorageDirectory(false),
		join.Empty(t.name, constant.LiteExtension),
	)
}
