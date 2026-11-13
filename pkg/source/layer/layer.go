package layer

import "github.com/funtimecoding/soil/pkg/source/snapshot"

type Layer struct {
	parent       *Layer
	roots        []string
	before       *snapshot.Snapshot
	written      map[string][]byte
	removed      map[string]bool
	original     map[string][]byte
	expected     map[string][]byte
	moves        map[string]string
	order        []string
	writes       int
	beforeCommit func()
}
