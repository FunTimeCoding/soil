package lint

import (
	"github.com/funtimecoding/soil/pkg/lint/concern"
	"github.com/funtimecoding/soil/pkg/lint/constant"
	"github.com/funtimecoding/soil/pkg/lint/option"
	"github.com/funtimecoding/soil/pkg/lint/output"
	"github.com/funtimecoding/soil/pkg/system"
	"os"
	"path/filepath"
)

func Lint(
	name string,
	root string,
	o *option.Lint,
) {
	repo, empty := Walk(root, o)
	configuration := o.Configuration

	if configuration != "" && !filepath.IsAbs(configuration) {
		configuration = repo.Absolute(configuration)
	}

	o.Registries = loadConfiguration(configuration).Registries
	r := output.NewResultsWithDirectory(repo.Root)
	Header(name, repo.Root, scopeDetail(o, repo.Files))

	for _, p := range empty {
		if !InScope(o, p) {
			continue
		}

		if o.Fix {
			system.Remove(repo.Absolute(p))
			r.AddConcern(
				concern.NewFile(
					"empty_directory",
					"removed empty directory",
					p,
					true,
				),
			)
		} else {
			r.AddConcern(
				concern.NewFile("empty_directory", "empty directory", p, false),
			)
		}
	}

	for _, p := range repo.Files.Files() {
		if !InScope(o, p) || Skipped(o, p) {
			continue
		}

		if repo.Files.FileAt(p).Size == 0 {
			if o.Fix {
				system.Remove(repo.Absolute(p))
				r.AddConcern(
					concern.NewFile("empty_file", "removed empty file", p, true),
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
				system.Remove(repo.Absolute(p))
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

	if o.Fix {
		fixes.Flush(repo.Root)
	}

	hasBlocked := output.PrintResults(r.Entries, o.Summary)

	if o.Census {
		output.PrintCensus(r.Unchecked)
	}

	if hasBlocked {
		os.Exit(1)
	}
}
