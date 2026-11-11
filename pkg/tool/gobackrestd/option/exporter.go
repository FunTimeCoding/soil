package option

import (
	"github.com/funtimecoding/soil/pkg/backrest/face"
	"github.com/funtimecoding/soil/pkg/web/option"
)

type Exporter struct {
	Metric  option.Web
	Verbose bool
	Source  func() face.BackrestClient
}
