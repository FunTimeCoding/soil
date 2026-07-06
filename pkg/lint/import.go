package lint

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/lint/constant"
	"github.com/funtimecoding/soil/pkg/lint/file_report"
	"github.com/funtimecoding/soil/pkg/strings/join"
	"github.com/funtimecoding/soil/pkg/system"
	"io"
	"slices"
	"strings"
)

func Import(
	path string,
	r io.Reader,
) *file_report.Report {
	s := file_report.New(path, r)
	var block []string
	var inside bool
	var start int
	var blockText string
	var reportedBlank bool

	for s.Scan() {
		line, number := s.Text()

		if !inside && strings.HasPrefix(line, "import (") {
			inside = true
			start = number
			blockText = line
		}

		if inside {
			if line == "" {
				if !reportedBlank {
					reportedBlank = true
					s.AddConcern(
						constant.ImportBlankKey,
						constant.ImportBlankText,
						path,
						start,
						blockText,
						true,
					)
				}

				continue
			}

			block = append(block, line)

			if strings.HasPrefix(line, ")") {
				inside = false

				if len(block) == 3 {
					s.ChangedLine(
						fmt.Sprintf(
							"import %s",
							strings.TrimSpace(block[1]),
						),
					)
					s.AddConcern(
						constant.SingleMultiImportKey,
						constant.SingleMultiImportText,
						path,
						start,
						blockText,
						true,
					)
				} else {
					sorted := make([]string, len(block)-2)
					copy(sorted, block[1:len(block)-1])
					SortImports(sorted)
					sorted = append([]string{"import ("}, sorted...)
					sorted = append(sorted, ")")

					if !slices.Equal(block, sorted) {
						s.AddConcern(
							constant.UnsortedImportsKey,
							constant.UnsortedImportsText,
							path,
							start,
							blockText,
							true,
						)
					}

					s.ChangedLine(join.NewLine(sorted))
				}

				block = block[:0]
			}
		} else {
			s.PassLine(line)
		}
	}

	s.Fix = func() {
		if s.Fixed != "" {
			fmt.Printf("Simplify import %s\n", path)
			system.SaveFile(path, s.Fixed)
		}
	}

	return s.Finalize()
}
