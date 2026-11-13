package overlay

import (
	"github.com/funtimecoding/soil/pkg/notation"
	"github.com/funtimecoding/soil/pkg/system"
	"github.com/funtimecoding/soil/pkg/tool/gobreak/constant"
	"path/filepath"
)

func (o *Overlay) Write() string {
	result := filepath.Join(o.directory, constant.OverlayFile)
	system.WriteFile(
		result,
		[]byte(notation.Encode(o, false)),
		constant.FileMode,
	)

	return result
}
