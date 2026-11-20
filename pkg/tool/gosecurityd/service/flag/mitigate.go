package flag

import (
	"github.com/funtimecoding/soil/pkg/errors/conflict"
	"github.com/funtimecoding/soil/pkg/tool/gosecurityd/store/flag"
	"github.com/funtimecoding/soil/pkg/vulnerability/identifier"
)

func (s *Service) Mitigate(
	name string,
	mitigation string,
) (*flag.Flag, error) {
	if e := identifier.Check(name); e != nil {
		return nil, e
	}

	result, f := s.store.ByVulnerability(name)

	if f != nil {
		return nil, f
	}

	if result.Mitigated {
		return nil, conflict.Format("flag already mitigated: %s", name)
	}

	result.Mitigated = true
	result.Mitigation = mitigation

	if g := s.store.Save(result); g != nil {
		return nil, g
	}

	return result, nil
}
