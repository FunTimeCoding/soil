package mock_client

import "github.com/funtimecoding/soil/pkg/tool/goproxmoxd/generated/client"

type Client struct {
	MachineList        []client.Machine
	ContainerList      []client.Container
	NodeList           []client.Node
	InstanceList       []client.InstanceDetail
	NetworkList        []client.Network
	StorageList        []client.Storage
	StorageContentList []client.StorageContentItem
	MachineSnapshots   map[int64][]client.Snapshot
	ContainerSnapshots map[int64][]client.Snapshot
	Snippets           map[string]string
	Plan               *client.ReconcilePlan
	Failure            error
	Closed             int
	tasks              int
}
