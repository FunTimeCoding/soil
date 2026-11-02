package pattern

import (
	"context"
	"github.com/funtimecoding/soil/pkg/lint/output"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/service/commit"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/sink"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/types/result"
)

func (s *Service) ApplyPattern(
	x context.Context,
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
	r, e := s.commit.Transact(
		x,
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

	if report != nil && r != nil && commit.HasUnfixed(r) {
		report.Applied = false
	}

	return r, report, e
}
