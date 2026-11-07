package option

import (
	"github.com/funtimecoding/soil/pkg/github/face"
	"github.com/funtimecoding/soil/pkg/web/option"
)

type Exporter struct {
	Metric  option.Web
	Owner   string
	Verbose bool
	Source  func() face.GitHubClient
}
