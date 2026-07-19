package lint

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/constant"
	"github.com/funtimecoding/soil/pkg/lint/concern"
	lintConstant "github.com/funtimecoding/soil/pkg/lint/constant"
	"github.com/funtimecoding/soil/pkg/lint/option"
	"github.com/funtimecoding/soil/pkg/lint/output"
	"github.com/funtimecoding/soil/pkg/system"
	"github.com/funtimecoding/soil/pkg/system/virtual_file_system"
	"os"
)

func Lint(
	skipString string,
	verbose bool,
	fix bool,
	summary bool,
) {
	skip := option.New(skipString, verbose)
	r := output.NewResults()

	for _, p := range system.EmptyDirectories(
		constant.CurrentDirectory,
		verbose,
	) {
		if Skipped(skip, p) {
			if verbose {
				fmt.Printf("Skip empty directory: %s\n", p)
			}

			continue
		}

		if fix {
			system.Remove(p)
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
				concern.NewFile(
					"empty_directory",
					"empty directory",
					p,
					false,
				),
			)
		}
	}

	for _, p := range system.FilesRecursive(
		constant.CurrentDirectory,
		verbose,
	) {
		if Skipped(skip, p) {
			if verbose {
				fmt.Printf("Skip file: %s\n", p)
			}

			continue
		}

		if !system.FileEmpty(p) {
			if verbose {
				fmt.Printf("Non empty file: %s\n", p)
			}

			continue
		}

		if fix {
			system.Remove(p)
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
				concern.NewFile(
					"empty_file",
					"empty file",
					p,
					false,
				),
			)
		}
	}

	v := virtual_file_system.From(constant.CurrentDirectory)

	for _, p := range v.Files() {
		if Skipped(skip, p) {
			continue
		}

		if isKnownBinaryExtension(p) {
			continue
		}

		if IsExecutable(v.ReadString(p)) {
			if fix {
				system.Remove(p)
				r.AddConcern(
					concern.NewFile(
						lintConstant.StrayBinaryKey,
						"removed stray binary",
						p,
						true,
					),
				)
			} else {
				r.AddConcern(
					concern.NewFile(
						lintConstant.StrayBinaryKey,
						lintConstant.StrayBinaryText,
						p,
						false,
					),
				)
			}
		}
	}

	fixes := Check(v, skip, fix, verbose, false, r)

	if fix {
		fixes.Flush(constant.CurrentDirectory)
	}

	hasBlocked := output.PrintResults(r.Entries, summary)

	if hasBlocked {
		os.Exit(1)
	}
}
