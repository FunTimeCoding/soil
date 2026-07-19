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

	all, fileSet := Load("", patterns)
	violations := FindViolations(all)

	if len(violations) == 0 {
		return
	}

	edits := BuildAllEdits(fileSet, all, violations, r)
	ApplyEdits(fileSet, edits, "", diff)

	if !diff {
		loadedFiles := BuildLoadedFiles(all)
		FixUnloadedReferences(violations, loadedFiles, "", r)
	}
}
