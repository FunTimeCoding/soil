package goreach

import (
	"github.com/funtimecoding/soil/pkg/lint/reach"
	"github.com/funtimecoding/soil/pkg/source/inventory"
	"github.com/funtimecoding/soil/pkg/source/warm_forest"
)

func Reach(
	trees *warm_forest.Forest,
	root string,
	transcript string,
) (*reach.Reach, *reach.Edits, []string, error) {
	modules := map[string]string{}

	for _, m := range inventory.LoadOptional(inventory.DefaultPath()).Modules {
		modules[m.Name] = m.Directory
	}

	edits, e := reach.ReadTranscript(transcript, modules)

	if e != nil {
		return nil, nil, nil, e
	}

	files, f := reach.ChangedEdits(root, edits)

	if f != nil {
		return nil, nil, nil, f
	}

	return reach.Directories(trees, root, files), edits, files, nil
}
