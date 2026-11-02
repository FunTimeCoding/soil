package face

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/goproxmoxd/generated/client"
)

type MachineClient interface {
	CloneMachine(
		x context.Context,
		identifier int64,
		p *client.CloneMachineParams,
		b client.CloneMachineJSONRequestBody,
	) (*client.CreateMachineResult, error)
	CreateMachine(
		x context.Context,
		p *client.CreateMachineParams,
		b client.CreateMachineJSONRequestBody,
	) (*client.CreateMachineResult, error)
	CreateMachineSnapshot(
		x context.Context,
		identifier int64,
		p *client.CreateMachineSnapshotParams,
		b client.CreateMachineSnapshotJSONRequestBody,
	) (*client.TaskResult, error)
	DeleteMachine(
		x context.Context,
		identifier int64,
		p *client.DeleteMachineParams,
	) (*client.TaskResult, error)
	DeleteMachineSnapshot(
		x context.Context,
		identifier int64,
		name string,
		p *client.DeleteMachineSnapshotParams,
	) (*client.TaskResult, error)
	GetMachine(
		x context.Context,
		identifier int64,
		p *client.GetMachineParams,
	) (*client.MachineDetail, error)
	ListMachineSnapshots(
		x context.Context,
		identifier int64,
		p *client.ListMachineSnapshotsParams,
	) ([]client.Snapshot, error)
	ListMachines(
		x context.Context,
		p *client.ListMachinesParams,
	) ([]client.Machine, error)
	ResetMachine(
		x context.Context,
		identifier int64,
		p *client.ResetMachineParams,
	) (*client.TaskResult, error)
	RollbackMachineSnapshot(
		x context.Context,
		identifier int64,
		name string,
		p *client.RollbackMachineSnapshotParams,
	) (*client.TaskResult, error)
	ShutdownMachine(
		x context.Context,
		identifier int64,
		p *client.ShutdownMachineParams,
	) (*client.TaskResult, error)
	StartMachine(
		x context.Context,
		identifier int64,
		p *client.StartMachineParams,
	) (*client.TaskResult, error)
	StopMachine(
		x context.Context,
		identifier int64,
		p *client.StopMachineParams,
	) (*client.TaskResult, error)
	UpdateMachine(
		x context.Context,
		identifier int64,
		p *client.UpdateMachineParams,
		b client.UpdateMachineJSONRequestBody,
	) (*client.StatusResult, error)
}
