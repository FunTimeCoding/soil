package option

import "github.com/funtimecoding/soil/pkg/tool/gogithubd/face"

type Exporter struct {
	MetricAddress string
	Owner         string
	Verbose       bool
	Source        func() face.RunSource
}
