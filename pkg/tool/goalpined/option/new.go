package option

import (
	"github.com/funtimecoding/soil/pkg/tool/goalpined/constant"
	"github.com/funtimecoding/soil/pkg/tool/goalpined/face"
	"github.com/funtimecoding/soil/pkg/tool/goalpined/indexer"
	"github.com/funtimecoding/soil/pkg/web/option"
)

func New() *Alpine {
	return &Alpine{
		Files:        option.Web{Address: constant.FileAddress},
		PackageRoot:  constant.PackageRoot,
		KeyDirectory: constant.KeyDirectory,
		Indexer:      func() face.Indexer { return indexer.New() },
	}
}
