package face

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/goproxmoxd/generated/client"
)

type NodeClient interface {
	DownloadLocator(
		x context.Context,
		node string,
		storage string,
		p *client.DownloadLocatorParams,
		b client.DownloadLocatorJSONRequestBody,
	) (*client.StatusResult, error)
	GetNodeStatus(
		x context.Context,
		node string,
		p *client.GetNodeStatusParams,
	) (*client.NodeStatus, error)
	ListNetworks(
		x context.Context,
		node string,
		p *client.ListNetworksParams,
	) ([]client.Network, error)
	ListNodes(
		x context.Context,
		p *client.ListNodesParams,
	) ([]client.Node, error)
	ListStorageContent(
		x context.Context,
		node string,
		storage string,
		p *client.ListStorageContentParams,
	) ([]client.StorageContentItem, error)
	ListStorages(
		x context.Context,
		node string,
		p *client.ListStoragesParams,
	) ([]client.Storage, error)
}
