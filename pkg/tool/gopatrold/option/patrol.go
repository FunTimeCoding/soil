package option

import (
	"github.com/funtimecoding/soil/pkg/patrol/face"
	"github.com/funtimecoding/soil/pkg/web/option"
)

type Patrol struct {
	Metric    option.Web
	Verbose   bool
	Instances []string
	Source    func(instance string) face.PatrolClient
}
