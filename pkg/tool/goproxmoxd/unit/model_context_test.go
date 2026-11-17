package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/tool/goproxmoxd/constant"
	"github.com/funtimecoding/soil/pkg/tool/goproxmoxd/unit/base"
	"testing"
)

func TestInstructionsOmitInstancesWithOne(t *testing.T) {
	o := base.New(t)
	c := o.Tools()
	assert.StringNotContains(
		t,
		"use_instance",
		c.InitializeResult().Instructions,
	)
}

func TestListNodes(t *testing.T) {
	o := base.New(t)
	c := o.Tools()
	result := c.MustCallTool(constant.ListNodes, nil)
	assert.StringContains(t, "alfa", result)
}

func TestListMachines(t *testing.T) {
	o := base.New(t)
	c := o.Tools()
	o.Alfa().Hypervisor.AddMachine("alfa", 100, "web-server")
	result := c.MustCallTool(
		constant.ListMachines,
		map[string]any{"node": "alfa"},
	)
	assert.StringContains(t, "web-server", result)
}

func TestCreateMachineCloudInit(t *testing.T) {
	o := base.New(t)
	c := o.Tools()
	result := c.MustCallTool(
		constant.CreateMachine,
		map[string]any{
			"node":    "alfa",
			"name":    "ci-test",
			"ci_user": "admin",
		},
	)
	assert.StringContains(t, "created", result)
}

func TestCreateMachineWithDiskImport(t *testing.T) {
	o := base.New(t)
	c := o.Tools()
	result := c.MustCallTool(
		constant.CreateMachine,
		map[string]any{
			"node":        "alfa",
			"name":        "import-test",
			"disk_import": "local:import/debian-13.qcow2",
			"disk_size":   64,
			"ci_user":     "admin",
		},
	)
	assert.StringContains(t, "created", result)
}

func TestCreateMachineCDROM(t *testing.T) {
	o := base.New(t)
	c := o.Tools()
	result := c.MustCallTool(
		constant.CreateMachine,
		map[string]any{
			"node":  "alfa",
			"name":  "iso-test",
			"cdrom": "local:iso/debian.iso",
		},
	)
	assert.StringContains(t, "created", result)
}

func TestCreateMachineCDROMAndCloudInitError(t *testing.T) {
	o := base.New(t)
	c := o.Tools()
	result := c.MustCallToolError(
		constant.CreateMachine,
		map[string]any{
			"node":    "alfa",
			"name":    "conflict",
			"cdrom":   "local:iso/debian.iso",
			"ci_user": "admin",
		},
	)
	assert.StringContains(t, "mutually exclusive", result)
}

func TestUpdateMachine(t *testing.T) {
	o := base.New(t)
	c := o.Tools()
	o.Alfa().Hypervisor.AddMachine("alfa", 100, "original")
	result := c.MustCallTool(
		constant.UpdateMachine,
		map[string]any{
			"identifier":  100,
			"node":        "alfa",
			"tags":        "web;prod",
			"description": "production web server",
		},
	)
	assert.StringContains(t, "updated", result)
}

func TestUpdateMachineDelete(t *testing.T) {
	o := base.New(t)
	c := o.Tools()
	o.Alfa().Hypervisor.AddMachine("alfa", 100, "tagged-vm")
	result := c.MustCallTool(
		constant.UpdateMachine,
		map[string]any{
			"identifier": 100,
			"node":       "alfa",
			"delete":     "tags,description",
		},
	)
	assert.StringContains(t, "updated", result)
}

func TestUpdateMachineSetAndDeleteConflict(t *testing.T) {
	o := base.New(t)
	c := o.Tools()
	o.Alfa().Hypervisor.AddMachine("alfa", 100, "conflict-vm")
	result := c.MustCallToolError(
		constant.UpdateMachine,
		map[string]any{
			"identifier": 100,
			"node":       "alfa",
			"tags":       "web",
			"delete":     "tags",
		},
	)
	assert.StringContains(t, "cannot set and delete", result)
}

func TestUpdateMachineNoChanges(t *testing.T) {
	o := base.New(t)
	c := o.Tools()
	o.Alfa().Hypervisor.AddMachine("alfa", 100, "original")
	result := c.MustCallToolError(
		constant.UpdateMachine,
		map[string]any{"identifier": 100, "node": "alfa"},
	)
	assert.StringContains(t, "no changes", result)
}

func TestCloneMachine(t *testing.T) {
	o := base.New(t)
	c := o.Tools()
	o.Alfa().Hypervisor.AddMachine("alfa", 100, "template")
	result := c.MustCallTool(
		constant.CloneMachine,
		map[string]any{
			"identifier": 100,
			"name":       "clone-of-template",
			"node":       "alfa",
		},
	)
	assert.StringContains(t, "cloned", result)
}

func TestSnippetCRUD(t *testing.T) {
	o := base.New(t)
	c := o.Tools()
	c.MustCallTool(
		constant.CreateSnippet,
		map[string]any{
			"name":    "test.yml",
			"content": "#cloud-config\npackages:\n  - vim\n",
		},
	)
	list := c.MustCallTool(constant.ListSnippets, nil)
	assert.StringContains(t, "test.yml", list)
	get := c.MustCallTool(
		constant.GetSnippet,
		map[string]any{"name": "test.yml"},
	)
	assert.StringContains(t, "cloud-config", get)
	c.MustCallTool(constant.DeleteSnippet, map[string]any{"name": "test.yml"})
	listAfter := c.MustCallTool(constant.ListSnippets, nil)
	assert.StringNotContains(t, "test.yml", listAfter)
}

func TestListStorages(t *testing.T) {
	o := base.New(t)
	c := o.Tools()
	o.Alfa().Hypervisor.AddStorage("alfa", "local", "dir", "iso,snippets")
	result := c.MustCallTool(
		constant.ListStorages,
		map[string]any{"node": "alfa"},
	)
	assert.StringContains(t, "local", result)
}

func TestGetMachine(t *testing.T) {
	o := base.New(t)
	c := o.Tools()
	o.Alfa().Hypervisor.AddMachine("alfa", 100, "web-server")
	result := c.MustCallTool(
		constant.GetMachine,
		map[string]any{"identifier": 100, "node": "alfa"},
	)
	assert.StringContains(t, "web-server", result)
}

func TestDeleteMachine(t *testing.T) {
	o := base.New(t)
	c := o.Tools()
	o.Alfa().Hypervisor.AddMachine("alfa", 100, "disposable")
	result := c.MustCallTool(
		constant.DeleteMachine,
		map[string]any{"identifier": 100, "node": "alfa"},
	)
	assert.StringContains(t, "deleted", result)
}

func TestDeleteMachineRunning(t *testing.T) {
	o := base.New(t)
	c := o.Tools()
	o.Alfa().Hypervisor.AddMachine("alfa", 100, "running-vm")
	o.Alfa().Hypervisor.SetMachineStatus("alfa", 100, "running")
	result := c.MustCallToolError(
		constant.DeleteMachine,
		map[string]any{"identifier": 100, "node": "alfa"},
	)
	assert.StringContains(t, "running", result)
}

func TestStartMachine(t *testing.T) {
	o := base.New(t)
	c := o.Tools()
	o.Alfa().Hypervisor.AddMachine("alfa", 100, "stopped-vm")
	result := c.MustCallTool(
		constant.StartMachine,
		map[string]any{"identifier": 100, "node": "alfa"},
	)
	assert.StringContains(t, "task_id", result)
}

func TestCreateMachineSnapshot(t *testing.T) {
	o := base.New(t)
	c := o.Tools()
	o.Alfa().Hypervisor.AddMachine("alfa", 100, "snapshot-target")
	result := c.MustCallTool(
		constant.CreateMachineSnapshot,
		map[string]any{
			"identifier": 100,
			"name":       "before-upgrade",
			"node":       "alfa",
		},
	)
	assert.StringContains(t, "task_id", result)
}

func TestMachineNotFound(t *testing.T) {
	o := base.New(t)
	c := o.Tools()
	result := c.MustCallToolError(
		constant.GetMachine,
		map[string]any{"identifier": 999, "node": "alfa"},
	)
	assert.StringContains(t, "machine not found: 999", result)
	assert.StringNotContains(t, "event_identifier", result)
}

func TestStartContainer(t *testing.T) {
	o := base.New(t)
	c := o.Tools()
	o.Alfa().Hypervisor.AddContainer("alfa", 101, "stopped-container")
	result := c.MustCallTool(
		constant.StartContainer,
		map[string]any{"identifier": 101, "node": "alfa"},
	)
	assert.StringContains(t, "task_id", result)
}

func TestStopContainer(t *testing.T) {
	o := base.New(t)
	c := o.Tools()
	o.Alfa().Hypervisor.AddContainer("alfa", 101, "running-container")
	result := c.MustCallTool(
		constant.StopContainer,
		map[string]any{"identifier": 101, "node": "alfa"},
	)
	assert.StringContains(t, "task_id", result)
}

func TestShutdownContainer(t *testing.T) {
	o := base.New(t)
	c := o.Tools()
	o.Alfa().Hypervisor.AddContainer("alfa", 101, "running-container")
	result := c.MustCallTool(
		constant.ShutdownContainer,
		map[string]any{"identifier": 101, "node": "alfa"},
	)
	assert.StringContains(t, "task_id", result)
}
