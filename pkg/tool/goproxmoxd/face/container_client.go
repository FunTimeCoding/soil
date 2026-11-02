package face

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/goproxmoxd/generated/client"
)

type ContainerClient interface {
	CreateContainerSnapshot(
		x context.Context,
		identifier int64,
		p *client.CreateContainerSnapshotParams,
		b client.CreateContainerSnapshotJSONRequestBody,
	) (*client.TaskResult, error)
	DeleteContainerSnapshot(
		x context.Context,
		identifier int64,
		name string,
		p *client.DeleteContainerSnapshotParams,
	) (*client.TaskResult, error)
	GetContainer(
		x context.Context,
		identifier int64,
		p *client.GetContainerParams,
	) (*client.ContainerDetail, error)
	ListContainerSnapshots(
		x context.Context,
		identifier int64,
		p *client.ListContainerSnapshotsParams,
	) ([]client.Snapshot, error)
	ListContainers(
		x context.Context,
		p *client.ListContainersParams,
	) ([]client.Container, error)
	RollbackContainerSnapshot(
		x context.Context,
		identifier int64,
		name string,
		p *client.RollbackContainerSnapshotParams,
	) (*client.TaskResult, error)
	ShutdownContainer(
		x context.Context,
		identifier int64,
		p *client.ShutdownContainerParams,
	) (*client.TaskResult, error)
	StartContainer(
		x context.Context,
		identifier int64,
		p *client.StartContainerParams,
	) (*client.TaskResult, error)
	StopContainer(
		x context.Context,
		identifier int64,
		p *client.StopContainerParams,
	) (*client.TaskResult, error)
}
