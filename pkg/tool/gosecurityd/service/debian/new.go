package debian

import "github.com/funtimecoding/soil/pkg/tool/gosecurityd/tracker_source"

func New(sources ...*tracker_source.Source) *Service {
	return &Service{sources: sources}
}
