package lint

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/lint/constant"
	"github.com/funtimecoding/soil/pkg/lint/file_report"
	"github.com/funtimecoding/soil/pkg/lint/pointer"
	"github.com/funtimecoding/soil/pkg/markup/front_matter"
	"github.com/funtimecoding/soil/pkg/strings/split"
	"io"
	"strings"
)

func Pointers(
	r *pointer.Resolver,
	unchecked func(string, int, string, constant.Reason),
) Checker {
	return func(
		path string,
		reader io.Reader,
	) *file_report.Report {
		b, e := io.ReadAll(reader)
		errors.PanicOnError(e)
		content := string(b)
		s := file_report.New(path, strings.NewReader(content))
		var bases []string
		dead := 0
		baseLine := 0
		skip := 0

		if f, found := front_matter.Extract(content); found {
			skip = f.Lines
			baseLine = f.Line(constant.BaseKey)
			var d declaration

			if f.Decode(&d) == nil {
				for _, declared := range d.Base {
					for _, value := range split.Comma(declared) {
						value = strings.TrimSpace(value)

						if value == "" {
							continue
						}

						if r.BaseExists(value) {
							bases = append(bases, value)
						} else {
							dead++
						}
					}
				}
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
					v := r.Resolve(path, bases, expanded)

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
