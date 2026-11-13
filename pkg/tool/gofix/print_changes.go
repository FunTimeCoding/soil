package gofix

import "github.com/funtimecoding/soil/pkg/source/types/layer_change"

func PrintChanges(changes []*layer_change.Change) {
	for _, c := range changes {
		printDiff(c.Path, c.Destination, c.Original, c.Current)
	}
}
