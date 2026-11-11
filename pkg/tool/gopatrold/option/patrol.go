package option

import (
	"github.com/funtimecoding/soil/pkg/patrol/face"
	"github.com/funtimecoding/soil/pkg/web/option"
)

type Patrol struct {
	Metric  option.Web
	Verbose bool
	Hosts   []string
	Source  func(host string) face.PatrolClient
}
