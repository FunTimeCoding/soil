package lint

import (
	"github.com/funtimecoding/soil/pkg/lint/constant"
	"github.com/funtimecoding/soil/pkg/lint/file_report"
	"github.com/funtimecoding/soil/pkg/lint/pointer"
	"io"
)

func Pointers(
	roots []string,
	exists func(string) bool,
	siblingExists func(string) bool,
) Checker {
	return func(
		path string,
		r io.Reader,
	) *file_report.Report {
		s := file_report.New(path, r)

		for s.Scan() {
			line, number := s.Text()
			s.PassLine(line)

			for _, extracted := range pointer.Extract(line) {
				for _, candidate := range pointer.Expand(extracted) {
					switch pointer.Classify(candidate, roots) {
					case pointer.Absolute:
						s.AddConcern(
							constant.AbsolutePointerKey,
							constant.AbsolutePointerText,
							path,
							number,
							line,
							false,
						)
					case pointer.Sibling:
						if siblingExists(pointer.Normalize(candidate)) {
							continue
						}

						relative, inside := pointer.Relative(path, candidate)

						if inside && exists(relative) {
							continue
						}

						s.AddConcern(
							constant.DeadPointerKey,
							constant.DeadPointerText,
							path,
							number,
							line,
							false,
						)
					case pointer.Repository:
						if !exists(pointer.Normalize(candidate)) {
							s.AddConcern(
								constant.DeadPointerKey,
								constant.DeadPointerText,
								path,
								number,
								line,
								false,
							)
						}
					}
				}
			}
		}

		return s.Finalize()
	}
}
