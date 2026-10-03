package gofix

import (
	"bytes"
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/lint/concern"
	"github.com/funtimecoding/soil/pkg/lint/output"
	"os"
)

func RunFormatFixWithDirectory(
	patterns []string,
	directory string,
	diff bool,
	r *output.Results,
) {
	if len(patterns) == 0 {
		patterns = []string{"./..."}
	}

	all, _ := Load(directory, patterns)

	for _, name := range formatFiles(all) {
		source, e := os.ReadFile(name)

		if e != nil {
			errors.Printf("read %s: %s\n", name, e)

			continue
		}

		f, e := formatFile(name, source)

		if e != nil {
			errors.Printf("format %s: %s\n", name, e)

			continue
		}

		if bytes.Equal(f.Source, source) {
			continue
		}

		for _, c := range f.Changes {
			r.AddConcern(
				concern.NewLine(c.Kind, c.Message, name, c.Line, "", !diff),
			)
		}

		if !f.Converged {
			errors.Printf("format fix did not converge for %s\n", name)
		}

		if diff {
			printDiff(name, source, f.Source)

			continue
		}

		if e = os.WriteFile(name, f.Source, 0644); e != nil {
			errors.Printf("write %s: %s\n", name, e)
		}
	}
}
