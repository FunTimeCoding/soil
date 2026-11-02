package face

import "context"

type ProxmoxClient interface {
	SnapshotClient
	MachineClient
	ContainerClient
	NodeClient
	AccessClient
	ClusterClient
	Reachable(x context.Context) error
	Close()
}
