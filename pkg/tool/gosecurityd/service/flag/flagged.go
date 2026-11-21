package flag

import (
	"github.com/funtimecoding/soil/pkg/tool/gosecurityd/store/flag"
	"github.com/funtimecoding/soil/pkg/tool/gosecurityd/subject"
)

func (s *Service) Flagged(name string) (*flag.Flag, error) {
	if e := subject.Check(name); e != nil {
		return nil, e
	}

	return s.store.BySubject(name)
}
