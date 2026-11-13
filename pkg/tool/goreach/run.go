package goreach

import (
	"github.com/funtimecoding/soil/pkg/source/warm_forest"
	"github.com/funtimecoding/soil/pkg/system"
	"github.com/funtimecoding/soil/pkg/system/writer"
	"os"
	"slices"
)

func Run(
	root string,
	transcript string,
	documents bool,
) {
	r, edits, files, e := Reach(warm_forest.New(), root, transcript)

	if e != nil {
		system.Exitf(1, "%s\n", e)
	}

	writer.Print(os.Stderr, "%s\n", Announcement(r, edits, files))
	result := r.Packages

	if documents {
		result = slices.Concat(result, r.Documents)
	}

	for _, d := range result {
		writer.Print(os.Stdout, "%s\n", d)
	}
}
