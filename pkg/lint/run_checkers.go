package lint

import (
	"github.com/funtimecoding/soil/pkg/lint/file_report"
	"github.com/funtimecoding/soil/pkg/lint/option"
	"github.com/funtimecoding/soil/pkg/lint/output"
	"github.com/funtimecoding/soil/pkg/lint/repository"
	"github.com/funtimecoding/soil/pkg/lint/types/checker_step"
	"github.com/funtimecoding/soil/pkg/source/layer"
	"github.com/funtimecoding/soil/pkg/system/writer"
	"strings"
)

func runCheckers(
	repo *repository.Repository,
	fixes *layer.Layer,
	paths []string,
	steps []*checker_step.Step,
	o *option.Lint,
	r *output.Results,
) {
	for _, p := range paths {
		if o.Verbose {
			writer.Print(o.Trace, "Process: %s\n", p)
		}

		original := repo.Files.ReadString(p)
		content := original

		for _, s := range steps {
			var result *file_report.Report

			if s.Pure {
				result = o.Documents.Result(
					s.Name,
					p,
					content,
					func() *file_report.Report {
						return s.Check(p, strings.NewReader(content))
					},
				)
			} else {
				result = s.Check(p, strings.NewReader(content))
			}

			for _, c := range result.Concerns {
				if !o.Fix {
					c.Fixed = false
				}

				r.AddConcern(c)
			}

			if result.Fixed != "" && result.Fixed != result.Original {
				content = result.Fixed
			}
		}

		if content != original {
			fixes.Rewrite(repo.Absolute(p), []byte(original), []byte(content))
		}
	}
}
