package gofix

import "github.com/funtimecoding/soil/pkg/console"

func printDiff(
	path string,
	original []byte,
	modified []byte,
) {
	if string(original) == string(modified) {
		return
	}

	console.Format("--- %s\n+++ %s\n", path, path)
	originalLines := splitLines(original)
	modifiedLines := splitLines(modified)

	for i := 0; i < len(originalLines) || i < len(modifiedLines); i++ {
		var o, m string

		if i < len(originalLines) {
			o = originalLines[i]
		}

		if i < len(modifiedLines) {
			m = modifiedLines[i]
		}

		if o != m {
			if o != "" {
				console.Format("-%s\n", o)
			}

			if m != "" {
				console.Format("+%s\n", m)
			}
		}
	}
}
