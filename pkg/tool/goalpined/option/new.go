package option

import (
	"github.com/funtimecoding/soil/pkg/tool/goalpined/constant"
	"github.com/funtimecoding/soil/pkg/tool/goalpined/face"
	"github.com/funtimecoding/soil/pkg/tool/goalpined/indexer"
)

func New() *Alpine {
	return &Alpine{
		PackageRoot:  constant.PackageRoot,
		KeyDirectory: constant.KeyDirectory,
		Indexer:      func() face.Indexer { return indexer.New() },
	}
}
