package gofix

import "github.com/funtimecoding/soil/pkg/lint/output"

func runImportAliasFix(
	patterns []string,
	diff bool,
	r *output.Results,
) {
	runImportAliasFixWithDirectory(patterns, "", diff, r)
}
