package debian

import "github.com/funtimecoding/soil/pkg/tool/gosecurityd/tracker_source"

type Service struct {
	sources []*tracker_source.Source
}
