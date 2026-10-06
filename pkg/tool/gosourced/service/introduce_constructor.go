package service

import (
	"github.com/funtimecoding/soil/pkg/lint/output"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/service/result"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/service/sink"
)

func (s *Service) IntroduceConstructor(
	directory string,
	packagePath string,
	name string,
	parameters []string,
	fromSites bool,
	assignRest bool,
	dryRun bool,
) (*output.Results, *result.Constructor, error) {
	var report *result.Constructor
	r, e := s.transact(
		directory,
		nil,
		dryRun,
		func(d string, out *sink.Sink) (*output.Results, error) {
			introduced, c, f := s.introduceConstructorThrough(
				d,
				packagePath,
				name,
				parameters,
				fromSites,
				assignRest,
				out,
			)
			report = c

			return introduced, f
		},
	)

	return r, report, e
}
