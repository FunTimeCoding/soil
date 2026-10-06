package gofix

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/source/resolve"
	"os"
)

func listReported(
	directory string,
	patterns []string,
) map[string]bool {
	result, e := resolve.ListPackages(directory, patterns...)

	if e != nil {
		errors.Printf("list: %s\n", e)
		os.Exit(1)
	}

	return result
}
