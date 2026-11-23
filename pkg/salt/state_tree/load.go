package state_tree

import (
	"github.com/funtimecoding/soil/pkg/salt/constant"
	"github.com/funtimecoding/soil/pkg/strings/join"
	"github.com/funtimecoding/soil/pkg/system"
	"slices"
	"strings"
)

func (t *Tree) load() {
	for _, d := range slices.Backward(t.directories) {
		t.loadRoot(d)
	}

	for name, relative := range t.states {
		t.parseState(
			name,
			strings.HasSuffix(
				relative,
				join.Empty("/", constant.InitName, constant.StateExtension),
			),
			string(system.ReadBytesUnsafe(t.files[relative])),
		)
	}

	if path, exists := t.files[constant.TopFile]; exists {
		t.parseTop(string(system.ReadBytesUnsafe(path)))
	}
}
