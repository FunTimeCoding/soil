package lint

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/lint/constant"
	"github.com/funtimecoding/soil/pkg/lint/file_report"
	"github.com/funtimecoding/soil/pkg/lint/pointer"
	"github.com/funtimecoding/soil/pkg/markup/front_matter"
	"io"
	"strings"
)

func Pointers(
	r *pointer.Resolver,
	unchecked func(
		string,
		int,
		string,
		constant.Reason,
	),
) Checker {
	return func(
		path string,
		reader io.Reader,
	) *file_report.Report {
		b, e := io.ReadAll(reader)
		errors.PanicOnError(e)
		content := string(b)
		s := file_report.New(path, strings.NewReader(content))
		declared := pointer.NewDeclared()
		dead := 0
		baseLine := 0
		skip := 0

		if f, found := front_matter.Extract(content); found {
			skip = f.Lines
			baseLine = f.Line(constant.BaseKey)
			var d declaration

			if f.Decode(&d) == nil {
				for _, value := range declaredValues(d.Base) {
					if r.BaseExists(value) {
						declared.Bases = append(declared.Bases, value)
					} else {
						dead++
					}
				}

				declared.Hosts = declaredValues(d.Hosts)
				declared.Commands = declaredValues(d.Commands)
			}
		}

		for s.Scan() {
			line, number := s.Text()
			s.PassLine(line)

			if number <= skip {
				if number == baseLine {
					for range dead {
						addPointerConcern(
							s,
							constant.VerdictDead,
							path,
							number,
							line,
						)
					}
				}

				continue
			}

			for _, c := range pointer.Extract(line) {
				for _, expanded := range pointer.Expand(c) {
					v := r.Resolve(path, declared, expanded)

					if v.Verdict == constant.VerdictTallied {
						unchecked(path, number, expanded.Span, v.Reason)

						continue
					}

					if v.Verdict != constant.VerdictLive {
						addPointerConcern(s, v.Verdict, path, number, line)
					}
				}
			}
		}

		return s.Finalize()
	}
}
