package overlay

import (
	"github.com/funtimecoding/soil/pkg/strings/join"
	"github.com/funtimecoding/soil/pkg/system"
	"github.com/funtimecoding/soil/pkg/tool/gobreak/constant"
	"path/filepath"
	"strconv"
)

func (o *Overlay) Add(
	path string,
	content string,
) {
	target := filepath.Join(
		o.directory,
		join.Empty(
			strconv.Itoa(len(o.Replace)),
			constant.NameSeparator,
			filepath.Base(path),
		),
	)
	system.WriteFile(target, []byte(content), constant.FileMode)
	o.Replace[path] = target
}
