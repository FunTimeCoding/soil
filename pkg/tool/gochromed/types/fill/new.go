package fill

import "github.com/funtimecoding/soil/pkg/tool/gochromed/types/selector"

func New(
	s *selector.Selector,
	uid string,
	value string,
) *Fill {
	return &Fill{Selector: s, UID: uid, Value: value, Snapshot: true}
}
