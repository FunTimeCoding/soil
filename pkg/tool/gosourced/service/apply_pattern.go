package service

import (
	"github.com/funtimecoding/soil/pkg/lint/output"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/service/result"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/service/sink"
)

func (s *Service) ApplyPattern(
	directory string,
	packagePath string,
	symbol string,
	receiver string,
	pattern string,
	replacement string,
	partial bool,
	dryRun bool,
) (*output.Results, *result.Apply, error) {
	var report *result.Apply
	r, e := s.transact(
		directory,
		nil,
		dryRun,
		func(d string, out *sink.Sink) (*output.Results, error) {
			applied, a, f := s.applyPatternThrough(
				d,
				packagePath,
				symbol,
				receiver,
				pattern,
				replacement,
				partial,
				out,
			)
			report = a

			return applied, f
		},
	)

	if report != nil && r != nil && hasUnfixed(r) {
		report.Applied = false
	}

	return r, report, e
}
