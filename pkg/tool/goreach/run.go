package goreach

import (
	"github.com/funtimecoding/soil/pkg/lint/reach"
	"github.com/funtimecoding/soil/pkg/source/inventory"
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
	modules := map[string]string{}

	for _, m := range inventory.LoadOptional(inventory.DefaultPath()).Modules {
		modules[m.Name] = m.Directory
	}

	edits, e := reach.ReadTranscript(transcript, modules)

	if e != nil {
		system.Exitf(1, "%s\n", e)
	}

	files, f := reach.ChangedEdits(root, edits)

	if f != nil {
		system.Exitf(1, "%s\n", f)
	}

	r := reach.Directories(root, files)
	writer.Print(
		os.Stderr,
		"goreach: %d changed files, %d package directories reached; Bash writes %d resolved, %d unresolved, %d fixer runs, %d scripts\n",
		len(files),
		len(r.Packages),
		edits.Resolved,
		edits.Unresolved,
		edits.Fixers,
		edits.Scripts,
	)
	result := r.Packages

	if documents {
		result = slices.Concat(result, r.Documents)
	}

	for _, d := range result {
		writer.Print(os.Stdout, "%s\n", d)
	}
}
