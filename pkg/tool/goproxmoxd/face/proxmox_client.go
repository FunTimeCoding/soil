package face

type ProxmoxClient interface {
	SnapshotClient
	MachineClient
	ContainerClient
	NodeClient
}
