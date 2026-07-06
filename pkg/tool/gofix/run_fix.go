package gofix

import "github.com/funtimecoding/soil/pkg/lint/output"

func runFix(
	patterns []string,
	diff bool,
	r *output.Results,
) {
	if len(patterns) == 0 {
		patterns = []string{"./..."}
	}

	all, fileSet := load("", patterns)
	violations := findViolations(all)

	if len(violations) == 0 {
		return
	}

	edits := buildAllEdits(fileSet, all, violations, r)
	applyEdits(fileSet, edits, "", diff)

	if !diff {
		loadedFiles := buildLoadedFiles(all)
		fixUnloadedReferences(violations, loadedFiles, "", r)
	}
}
