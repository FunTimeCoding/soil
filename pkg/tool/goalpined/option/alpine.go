package option

import (
	"github.com/funtimecoding/soil/pkg/tool/goalpined/face"
	"github.com/funtimecoding/soil/pkg/web/option"
)

type Alpine struct {
	option.Web
	Files        option.Web
	PackageRoot  string
	KeyDirectory string
	SignatureKey string
	Indexer      func() face.Indexer
}
