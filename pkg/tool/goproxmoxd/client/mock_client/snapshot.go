package mock_client

import (
	"github.com/funtimecoding/soil/pkg/errors/not_found"
	"github.com/funtimecoding/soil/pkg/tool/goproxmoxd/generated/client"
)

func snapshot(
	snapshots []client.Snapshot,
	name string,
) (int, error) {
	for i, s := range snapshots {
		if s.Name == name {
			return i, nil
		}
	}

	return 0, not_found.Format("snapshot %s", name)
}
