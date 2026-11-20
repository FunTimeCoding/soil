package flag

import (
	"context"
	"github.com/funtimecoding/soil/pkg/errors/conflict"
	"github.com/funtimecoding/soil/pkg/errors/not_found"
	"github.com/funtimecoding/soil/pkg/tool/gosecurityd/constant"
	"github.com/funtimecoding/soil/pkg/tool/gosecurityd/store/flag"
	"github.com/funtimecoding/soil/pkg/vulnerability/identifier"
)

func (s *Service) Flag(
	x context.Context,
	name string,
	note string,
) (*flag.Flag, error) {
	if e := identifier.Check(name); e != nil {
		return nil, e
	}

	result, f := s.store.ByVulnerability(name)

	if f != nil && !not_found.Is(f) {
		return nil, f
	}

	if f == nil && !result.Mitigated {
		return nil, conflict.Exists(constant.FlagKind, name)
	}

	if f != nil {
		result = flag.Stub()
		result.Vulnerability = name
	}

	v, g := s.vulnerability.Vulnerability(x, name)

	if g != nil && !not_found.Is(g) {
		return nil, g
	}

	if g == nil {
		result.Severity = v.Severity
		result.Score = v.Score
	}

	result.Note = note
	result.Mitigated = false
	result.Mitigation = ""
	write := s.store.Save

	if f != nil {
		write = s.store.Create
	}

	if h := write(result); h != nil {
		return nil, h
	}

	return result, nil
}
