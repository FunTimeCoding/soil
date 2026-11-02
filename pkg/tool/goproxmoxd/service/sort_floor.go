package service

import (
	"github.com/funtimecoding/soil/pkg/tool/goproxmoxd/types/floor"
	"github.com/funtimecoding/soil/pkg/tool/goproxmoxd/types/floor/guest"
	"github.com/funtimecoding/soil/pkg/tool/goproxmoxd/types/floor/node"
	"github.com/funtimecoding/soil/pkg/tool/goproxmoxd/types/floor/storage"
	"slices"
	"strings"
)

func sortFloor(f *floor.Floor) {
	slices.SortFunc(
		f.Nodes,
		func(
			a node.Node,
			b node.Node,
		) int {
			if c := strings.Compare(a.Hypervisor, b.Hypervisor); c != 0 {
				return c
			}

			return strings.Compare(a.Name, b.Name)
		},
	)
	slices.SortFunc(
		f.Guests,
		func(
			a guest.Guest,
			b guest.Guest,
		) int {
			if c := strings.Compare(a.Hypervisor, b.Hypervisor); c != 0 {
				return c
			}

			if c := strings.Compare(a.Node, b.Node); c != 0 {
				return c
			}

			return strings.Compare(a.Name, b.Name)
		},
	)
	slices.SortFunc(
		f.Storages,
		func(
			a storage.Storage,
			b storage.Storage,
		) int {
			if c := strings.Compare(a.Hypervisor, b.Hypervisor); c != 0 {
				return c
			}

			return strings.Compare(a.Name, b.Name)
		},
	)
}
