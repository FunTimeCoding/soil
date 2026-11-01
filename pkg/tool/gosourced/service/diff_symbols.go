package service

import (
	"github.com/funtimecoding/soil/pkg/errors/not_found"
	gitConstant "github.com/funtimecoding/soil/pkg/git/constant"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/constant"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/service/report"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/service/revision"
)

func (s *Service) DiffSymbols(
	name string,
	from string,
	to string,
) (*report.Report, error) {
	m, okay := s.Module(name)

	if !okay {
		return nil, not_found.New(constant.ModuleKind, name)
	}

	if from == "" {
		from = gitConstant.HeadReference
	}

	before, e := revisionFiles(m.Directory, from)

	if e != nil {
		return nil, e
	}

	after, f := revisionFiles(m.Directory, to)

	if f != nil {
		return nil, f
	}

	return report.New(revision.New(before), revision.New(after)), nil
}
