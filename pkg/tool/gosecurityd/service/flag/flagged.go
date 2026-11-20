package flag

import (
	"github.com/funtimecoding/soil/pkg/tool/gosecurityd/store/flag"
	"github.com/funtimecoding/soil/pkg/vulnerability/identifier"
)

func (s *Service) Flagged(name string) (*flag.Flag, error) {
	if e := identifier.Check(name); e != nil {
		return nil, e
	}

	return s.store.ByVulnerability(name)
}
