package goproxmox

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/goproxmox/command_context"
	"github.com/funtimecoding/soil/pkg/tool/goproxmoxd/generated/client"
)

func runCreateMachine(c *command_context.Context, f *machineFlags) {
	body := client.CreateMachineJSONRequestBody{Node: f.node, Name: f.name}

	if f.diskImport != "" {
		body.DiskImport = &f.diskImport
	}

	if f.diskSize > 0 {
		body.DiskSize = &f.diskSize
	}

	if f.diskStorage != "" {
		body.DiskStorage = &f.diskStorage
	}

	if f.cdrom != "" {
		body.Cdrom = &f.cdrom
	}

	if f.bridge != "" {
		body.Bridge = &f.bridge
	}

	if f.cpuType != "" {
		body.CpuType = &f.cpuType
	}

	if f.operatingSystemType != "" {
		body.OsType = &f.operatingSystemType
	}

	if f.ciUser != "" {
		body.CiUser = &f.ciUser
	}

	if f.ciPassword != "" {
		body.CiPassword = &f.ciPassword
	}

	if f.sshKeys != "" {
		body.SshKeys = &f.sshKeys
	}

	if f.ipConfiguration != "" {
		body.IpConfig = &f.ipConfiguration
	}

	if f.searchDomain != "" {
		body.SearchDomain = &f.searchDomain
	}

	if f.cores > 0 {
		body.Cores = &f.cores
	}

	if f.memory > 0 {
		body.Memory = &f.memory
	}

	if f.tags != "" {
		body.Tags = &f.tags
	}

	if f.extras != "" {
		body.Extras = &f.extras
	}

	if f.start {
		body.Start = &f.start
	}

	c.Terminal().EmitValue(
		c.Client().CreateMachine(
			context.Background(),
			&client.CreateMachineParams{Instance: c.Instance()},
			body,
		),
	)
}
