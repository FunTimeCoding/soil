package lint

import (
	"github.com/funtimecoding/soil/pkg/lint/concern"
	"github.com/funtimecoding/soil/pkg/lint/constant"
	"github.com/funtimecoding/soil/pkg/lint/installed"
	"github.com/funtimecoding/soil/pkg/lint/option"
	"github.com/funtimecoding/soil/pkg/lint/output"
)

func Report(
	name string,
	root string,
	o *option.Lint,
) *output.Results {
	repo, empty := Walk(root, o)
	c := loadConfiguration(repo.Root)
	o.Registries = c.Registries
	o.ReflowSkips = c.Reflow.Skip
	r := output.NewResultsWithDirectory(repo.Root)
	r.Header = HeaderLine(name, repo.Root, scopeDetail(o, repo.Files))

	for _, p := range empty {
		if !InScope(o, p) {
			continue
		}

		if !o.Fix {
			r.AddConcern(
				concern.NewFile("empty_directory", "empty directory", p, false),
			)

			continue
		}

		if removeEmptyDirectory(repo.Absolute(p)) {
			r.AddConcern(
				concern.NewFile(
					"empty_directory",
					"removed empty directory",
					p,
					true,
				),
			)
		}
	}

	var removals []string

	for _, p := range repo.Files.Files() {
		if !InScope(o, p) || Skipped(o.Skips, p) {
			continue
		}

		if repo.Files.FileAt(p).Size == 0 {
			if o.Fix {
				removals = append(removals, p)
				r.AddConcern(
					concern.NewFile(
						"empty_file",
						"removed empty file",
						p,
						true,
					),
				)
			} else {
				r.AddConcern(
					concern.NewFile("empty_file", "empty file", p, false),
				)
			}

			continue
		}

		if isKnownBinaryExtension(p) {
			continue
		}

		if IsExecutable(repo.Files.ReadString(p)) {
			if o.Fix {
				removals = append(removals, p)
				r.AddConcern(
					concern.NewFile(
						constant.StrayBinaryKey,
						"removed stray binary",
						p,
						true,
					),
				)
			} else {
				r.AddConcern(
					concern.NewFile(
						constant.StrayBinaryKey,
						constant.StrayBinaryText,
						p,
						false,
					),
				)
			}
		}
	}

	fixes := Check(repo, o, r)

	for _, p := range removals {
		path := repo.Absolute(p)
		fixes.Guard(path, []byte(repo.Files.ReadString(p)))
		fixes.Remove(path)
	}

	if o.Fix && o.Layer == nil {
		commitFixes(fixes, o, r)
	}

	if len(o.Scopes) == 0 {
		for _, c := range installed.Check(repo.Root) {
			r.AddConcern(c)
		}
	}

	return r
}
