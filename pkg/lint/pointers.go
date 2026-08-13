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
	ignored func(string) bool,
) Checker {
	return func(
		path string,
		r io.Reader,
	) *file_report.Report {
		s := file_report.New(path, r)

		for s.Scan() {
			line, number := s.Text()
			s.PassLine(line)

			for _, bare := range pointer.ExtractBareLinks(line) {
				relative, inside := pointer.Relative(path, bare)

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
			}

			for _, extracted := range pointer.Extract(line) {
				for _, candidate := range pointer.Expand(extracted) {
					switch pointer.Classify(candidate, roots) {
					case constant.PointerAbsolute:
						s.AddConcern(
							constant.AbsolutePointerKey,
							constant.AbsolutePointerText,
							path,
							number,
							line,
							false,
						)
					case constant.PointerSibling:
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
					case constant.PointerRepository:
						normalized := pointer.Normalize(candidate)

						if exists(normalized) || ignored(normalized) {
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
					}
				}
			}
		}

		return s.Finalize()
	}
}
