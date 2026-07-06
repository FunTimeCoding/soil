package model_context

import (
	"github.com/funtimecoding/soil/pkg/tool/goproxmoxd/constant"
	"github.com/mark3labs/mcp-go/mcp"
)

func (s *Server) register() {
	if len(s.service.Instances()) > 1 {
		s.server.AddTool(
			mcp.NewTool(
				constant.ListInstances,
				mcp.WithDescription(
					"List all configured Proxmox instances. Shows which instance is currently active.",
				),
			),
			mcp.NewTypedToolHandler(s.listInstances),
		)
		s.server.AddTool(
			mcp.NewTool(
				constant.UseInstance,
				mcp.WithDescription(
					"Set the active Proxmox instance for this session. Required before using any other tool.",
				),
				mcp.WithString(
					"instance",
					mcp.Required(),
					mcp.Description("Instance name from list_instances"),
				),
			),
			mcp.NewTypedToolHandler(s.useInstance),
		)
	}

	s.server.AddTool(
		mcp.NewTool(
			constant.ListNodes,
			mcp.WithDescription("List all Proxmox nodes with status"),
		),
		mcp.NewTypedToolHandler(s.ListNodes),
	)
	s.server.AddTool(
		mcp.NewTool(
			constant.GetNodeStatus,
			mcp.WithDescription("Get detailed resource status for a Proxmox node"),
			mcp.WithString(
				"node",
				mcp.Required(),
				mcp.Description("Node name"),
			),
		),
		mcp.NewTypedToolHandler(s.GetNodeStatus),
	)
	s.server.AddTool(
		mcp.NewTool(
			constant.ListMachines,
			mcp.WithDescription(
				"List virtual machines, optionally filtered by node",
			),
			mcp.WithString(
				"node",
				mcp.Description("Node name (omit to list across all nodes)"),
			),
		),
		mcp.NewTypedToolHandler(s.ListMachines),
	)
	s.server.AddTool(
		mcp.NewTool(
			constant.GetMachine,
			mcp.WithDescription(
				"Get status and config for a virtual machine by VMID",
			),
			mcp.WithNumber(
				"identifier",
				mcp.Required(),
				mcp.Description("VM ID"),
			),
			mcp.WithString(
				"node",
				mcp.Description("Node name (omit to search all nodes)"),
			),
		),
		mcp.NewTypedToolHandler(s.GetMachine),
	)
	s.server.AddTool(
		mcp.NewTool(
			constant.StartMachine,
			mcp.WithDescription("Start a virtual machine"),
			mcp.WithNumber(
				"identifier",
				mcp.Required(),
				mcp.Description("Machine ID"),
			),
			mcp.WithString(
				"node",
				mcp.Description("Node name (omit to search all nodes)"),
			),
		),
		mcp.NewTypedToolHandler(s.StartMachine),
	)
	s.server.AddTool(
		mcp.NewTool(
			constant.StopMachine,
			mcp.WithDescription("Stop a virtual machine immediately"),
			mcp.WithNumber(
				"identifier",
				mcp.Required(),
				mcp.Description("Machine ID"),
			),
			mcp.WithString(
				"node",
				mcp.Description("Node name (omit to search all nodes)"),
			),
		),
		mcp.NewTypedToolHandler(s.StopMachine),
	)
	s.server.AddTool(
		mcp.NewTool(
			constant.ShutdownMachine,
			mcp.WithDescription("Gracefully shutdown a virtual machine"),
			mcp.WithNumber(
				"identifier",
				mcp.Required(),
				mcp.Description("Machine ID"),
			),
			mcp.WithString(
				"node",
				mcp.Description("Node name (omit to search all nodes)"),
			),
		),
		mcp.NewTypedToolHandler(s.ShutdownMachine),
	)
	s.server.AddTool(
		mcp.NewTool(
			constant.ResetMachine,
			mcp.WithDescription("Reset a virtual machine"),
			mcp.WithNumber(
				"identifier",
				mcp.Required(),
				mcp.Description("Machine ID"),
			),
			mcp.WithString(
				"node",
				mcp.Description("Node name (omit to search all nodes)"),
			),
		),
		mcp.NewTypedToolHandler(s.ResetMachine),
	)
	s.server.AddTool(
		mcp.NewTool(
			constant.ListMachineSnapshots,
			mcp.WithDescription("List snapshots for a virtual machine"),
			mcp.WithNumber(
				"identifier",
				mcp.Required(),
				mcp.Description("Machine ID"),
			),
			mcp.WithString(
				"node",
				mcp.Description("Node name (omit to search all nodes)"),
			),
		),
		mcp.NewTypedToolHandler(s.ListMachineSnapshots),
	)
	s.server.AddTool(
		mcp.NewTool(
			constant.CreateMachineSnapshot,
			mcp.WithDescription("Create a snapshot of a virtual machine"),
			mcp.WithNumber(
				"identifier",
				mcp.Required(),
				mcp.Description("Machine ID"),
			),
			mcp.WithString(
				"name",
				mcp.Required(),
				mcp.Description("Snapshot name"),
			),
			mcp.WithString(
				"node",
				mcp.Description("Node name (omit to search all nodes)"),
			),
		),
		mcp.NewTypedToolHandler(s.CreateMachineSnapshot),
	)
	s.server.AddTool(
		mcp.NewTool(
			constant.RollbackMachineSnapshot,
			mcp.WithDescription("Rollback a virtual machine to a snapshot"),
			mcp.WithNumber(
				"identifier",
				mcp.Required(),
				mcp.Description("Machine ID"),
			),
			mcp.WithString(
				"name",
				mcp.Required(),
				mcp.Description("Snapshot name"),
			),
			mcp.WithString(
				"node",
				mcp.Description("Node name (omit to search all nodes)"),
			),
		),
		mcp.NewTypedToolHandler(s.RollbackMachineSnapshot),
	)
	s.server.AddTool(
		mcp.NewTool(
			constant.DeleteMachineSnapshot,
			mcp.WithDescription("Delete a snapshot from a virtual machine"),
			mcp.WithNumber(
				"identifier",
				mcp.Required(),
				mcp.Description("Machine ID"),
			),
			mcp.WithString(
				"name",
				mcp.Required(),
				mcp.Description("Snapshot name"),
			),
			mcp.WithString(
				"node",
				mcp.Description("Node name (omit to search all nodes)"),
			),
		),
		mcp.NewTypedToolHandler(s.DeleteMachineSnapshot),
	)
	s.server.AddTool(
		mcp.NewTool(
			constant.CreateMachine,
			mcp.WithDescription(
				"Create a virtual machine. Waits for the creation task to complete before returning.",
			),
			mcp.WithString(
				"node",
				mcp.Required(),
				mcp.Description("Target node name"),
			),
			mcp.WithString(
				"name",
				mcp.Required(),
				mcp.Description("VM name"),
			),
			mcp.WithNumber(
				"identifier",
				mcp.Description("VMID (auto-allocated when omitted)"),
			),
			mcp.WithNumber("cores", mcp.Description("CPU cores (default 2)")),
			mcp.WithNumber(
				"memory",
				mcp.Description("Memory in MiB (default 2048)"),
			),
			mcp.WithNumber(
				"disk_size",
				mcp.Description("Disk size in GiB (default 32)"),
			),
			mcp.WithString(
				"disk_storage",
				mcp.Description("Storage backend (default local-lvm)"),
			),
			mcp.WithString(
				"disk_import",
				mcp.Description("Import-from path for disk image, e.g. local:import/debian-13-generic-amd64.qcow2"),
			),
			mcp.WithString(
				"cdrom",
				mcp.Description("ISO volume to mount as CD-ROM, e.g. local:iso/debian-13.iso. Boots from CD-ROM first when set."),
			),
			mcp.WithString(
				"bridge",
				mcp.Description("Network bridge (default vmbr0)"),
			),
			mcp.WithString(
				"cpu_type",
				mcp.Description("CPU type (default host)"),
			),
			mcp.WithBoolean(
				"agent",
				mcp.Description("Enable QEMU guest agent (default true)"),
			),
			mcp.WithBoolean(
				"onboot",
				mcp.Description("Start VM automatically on host boot"),
			),
			mcp.WithString("ci_user", mcp.Description("Cloud-init user")),
			mcp.WithString(
				"ci_password",
				mcp.Description("Cloud-init password"),
			),
			mcp.WithString(
				"ssh_keys",
				mcp.Description("SSH public keys, newline-separated"),
			),
			mcp.WithString(
				"ip_config",
				mcp.Description("Cloud-init IP config (default ip=dhcp when cloud-init is used)"),
			),
			mcp.WithString(
				"search_domain",
				mcp.Description("DNS search domain for cloud-init, e.g. local"),
			),
			mcp.WithString(
				"os_type",
				mcp.Description("OS type, e.g. l26 for Linux"),
			),
			mcp.WithBoolean(
				"start",
				mcp.Description("Start VM after creation (default false)"),
			),
			mcp.WithString(
				"tags",
				mcp.Description("Semicolon-separated tags"),
			),
			mcp.WithString(
				"extras",
				mcp.Description("Additional Proxmox options as comma-separated key=value pairs"),
			),
		),
		mcp.NewTypedToolHandler(s.CreateMachine),
	)
	s.server.AddTool(
		mcp.NewTool(
			constant.UpdateMachine,
			mcp.WithDescription(
				"Update configuration of an existing virtual machine.",
			),
			mcp.WithNumber(
				"identifier",
				mcp.Required(),
				mcp.Description("Machine ID"),
			),
			mcp.WithString(
				"node",
				mcp.Description("Node name (omit to search all nodes)"),
			),
			mcp.WithString("name", mcp.Description("New VM name")),
			mcp.WithString(
				"tags",
				mcp.Description("Semicolon-separated tags (replaces existing)"),
			),
			mcp.WithBoolean(
				"onboot",
				mcp.Description("Start VM automatically on host boot"),
			),
			mcp.WithNumber("cores", mcp.Description("CPU cores")),
			mcp.WithNumber("memory", mcp.Description("Memory in MiB")),
			mcp.WithString(
				"description",
				mcp.Description("VM description"),
			),
			mcp.WithString(
				"delete",
				mcp.Description("Comma-separated field names to clear, e.g. tags,description"),
			),
		),
		mcp.NewTypedToolHandler(s.UpdateMachine),
	)
	s.server.AddTool(
		mcp.NewTool(
			constant.CloneMachine,
			mcp.WithDescription(
				"Clone a virtual machine. Waits for the clone task to complete before returning.",
			),
			mcp.WithNumber(
				"identifier",
				mcp.Required(),
				mcp.Description("Source machine ID to clone from"),
			),
			mcp.WithString(
				"name",
				mcp.Required(),
				mcp.Description("Name for the new VM"),
			),
			mcp.WithString(
				"node",
				mcp.Description("Node name (omit to search all nodes)"),
			),
			mcp.WithNumber(
				"new_identifier",
				mcp.Description("VMID for the clone (auto-allocated when omitted)"),
			),
			mcp.WithBoolean(
				"full",
				mcp.Description("Full clone instead of linked clone (default false)"),
			),
			mcp.WithString(
				"storage",
				mcp.Description("Target storage for the clone"),
			),
			mcp.WithString(
				"snapshot",
				mcp.Description("Clone from a specific snapshot instead of current state"),
			),
		),
		mcp.NewTypedToolHandler(s.CloneMachine),
	)
	s.server.AddTool(
		mcp.NewTool(
			constant.DeleteMachine,
			mcp.WithDescription(
				"Delete a virtual machine. The VM must be stopped first. Waits for deletion to complete.",
			),
			mcp.WithNumber(
				"identifier",
				mcp.Required(),
				mcp.Description("Machine ID"),
			),
			mcp.WithString(
				"node",
				mcp.Description("Node name (omit to search all nodes)"),
			),
			mcp.WithBoolean(
				"purge",
				mcp.Description("Remove from cluster config and associated resources (default false)"),
			),
		),
		mcp.NewTypedToolHandler(s.DeleteMachine),
	)
	s.server.AddTool(
		mcp.NewTool(
			constant.ListContainers,
			mcp.WithDescription(
				"List LXC containers, optionally filtered by node",
			),
			mcp.WithString(
				"node",
				mcp.Description("Node name (omit to list across all nodes)"),
			),
		),
		mcp.NewTypedToolHandler(s.ListContainers),
	)
	s.server.AddTool(
		mcp.NewTool(
			constant.GetContainer,
			mcp.WithDescription(
				"Get status and config for an LXC container by VMID",
			),
			mcp.WithNumber(
				"identifier",
				mcp.Required(),
				mcp.Description("Container ID"),
			),
			mcp.WithString(
				"node",
				mcp.Description("Node name (omit to search all nodes)"),
			),
		),
		mcp.NewTypedToolHandler(s.GetContainer),
	)
	s.server.AddTool(
		mcp.NewTool(
			constant.ListContainerSnapshots,
			mcp.WithDescription("List snapshots for an LXC container"),
			mcp.WithNumber(
				"identifier",
				mcp.Required(),
				mcp.Description("Container ID"),
			),
			mcp.WithString(
				"node",
				mcp.Description("Node name (omit to search all nodes)"),
			),
		),
		mcp.NewTypedToolHandler(s.ListContainerSnapshots),
	)
	s.server.AddTool(
		mcp.NewTool(
			constant.CreateContainerSnapshot,
			mcp.WithDescription("Create a snapshot of an LXC container"),
			mcp.WithNumber(
				"identifier",
				mcp.Required(),
				mcp.Description("Container ID"),
			),
			mcp.WithString(
				"name",
				mcp.Required(),
				mcp.Description("Snapshot name"),
			),
			mcp.WithString(
				"node",
				mcp.Description("Node name (omit to search all nodes)"),
			),
		),
		mcp.NewTypedToolHandler(s.CreateContainerSnapshot),
	)
	s.server.AddTool(
		mcp.NewTool(
			constant.RollbackContainerSnapshot,
			mcp.WithDescription("Rollback an LXC container to a snapshot"),
			mcp.WithNumber(
				"identifier",
				mcp.Required(),
				mcp.Description("Container ID"),
			),
			mcp.WithString(
				"name",
				mcp.Required(),
				mcp.Description("Snapshot name"),
			),
			mcp.WithString(
				"node",
				mcp.Description("Node name (omit to search all nodes)"),
			),
		),
		mcp.NewTypedToolHandler(s.RollbackContainerSnapshot),
	)
	s.server.AddTool(
		mcp.NewTool(
			constant.DeleteContainerSnapshot,
			mcp.WithDescription("Delete a snapshot from an LXC container"),
			mcp.WithNumber(
				"identifier",
				mcp.Required(),
				mcp.Description("Container ID"),
			),
			mcp.WithString(
				"name",
				mcp.Required(),
				mcp.Description("Snapshot name"),
			),
			mcp.WithString(
				"node",
				mcp.Description("Node name (omit to search all nodes)"),
			),
		),
		mcp.NewTypedToolHandler(s.DeleteContainerSnapshot),
	)
	s.server.AddTool(
		mcp.NewTool(
			constant.ListNetworks,
			mcp.WithDescription(
				"List network interfaces on a Proxmox node",
			),
			mcp.WithString(
				"node",
				mcp.Required(),
				mcp.Description("Node name"),
			),
		),
		mcp.NewTypedToolHandler(s.ListNetworks),
	)
	s.server.AddTool(
		mcp.NewTool(
			constant.ListStorages,
			mcp.WithDescription(
				"List storage backends on a node. Shows type, content types, and usage.",
			),
			mcp.WithString(
				"node",
				mcp.Required(),
				mcp.Description("Node name"),
			),
		),
		mcp.NewTypedToolHandler(s.ListStorages),
	)
	s.server.AddTool(
		mcp.NewTool(
			constant.ListStorageContent,
			mcp.WithDescription(
				"List content on a storage backend. Returns ISOs, disk images, templates, and backups with their volume identifiers.",
			),
			mcp.WithString(
				"node",
				mcp.Required(),
				mcp.Description("Node name"),
			),
			mcp.WithString(
				"storage",
				mcp.Required(),
				mcp.Description("Storage name from list_storages"),
			),
		),
		mcp.NewTypedToolHandler(s.ListStorageContent),
	)
	s.server.AddTool(
		mcp.NewTool(
			constant.DownloadLocator,
			mcp.WithDescription(
				"Download a file from a URL to a storage backend. Use content 'import' for qcow2 cloud images, 'iso' for ISOs, 'vztmpl' for container templates. Waits for the download to complete.",
			),
			mcp.WithString(
				"node",
				mcp.Required(),
				mcp.Description("Node name"),
			),
			mcp.WithString(
				"storage",
				mcp.Required(),
				mcp.Description("Storage name from list_storages"),
			),
			mcp.WithString(
				"content",
				mcp.Required(),
				mcp.Description("Content type: import, iso, or vztmpl"),
			),
			mcp.WithString(
				"filename",
				mcp.Required(),
				mcp.Description("Filename to save as on the storage"),
			),
			mcp.WithString(
				"locator",
				mcp.Required(),
				mcp.Description("URL to download from"),
			),
		),
		mcp.NewTypedToolHandler(s.DownloadLocator),
	)
	s.server.AddTool(
		mcp.NewTool(
			constant.CreateSnippet,
			mcp.WithDescription(
				"Create a cloud-init or other snippet file on the Proxmox host via SSH. Returns the volume reference for use with cicustom.",
			),
			mcp.WithString(
				"name",
				mcp.Required(),
				mcp.Description("Snippet filename, e.g. cloud-init-agent.yml"),
			),
			mcp.WithString(
				"content",
				mcp.Required(),
				mcp.Description("File content, e.g. a cloud-init YAML body"),
			),
		),
		mcp.NewTypedToolHandler(s.CreateSnippet),
	)
	s.server.AddTool(
		mcp.NewTool(
			constant.GetSnippet,
			mcp.WithDescription("Read a snippet file from the Proxmox host"),
			mcp.WithString(
				"name",
				mcp.Required(),
				mcp.Description("Snippet filename"),
			),
		),
		mcp.NewTypedToolHandler(s.GetSnippet),
	)
	s.server.AddTool(
		mcp.NewTool(
			constant.ListSnippets,
			mcp.WithDescription(
				"List snippet files on the Proxmox host",
			),
		),
		mcp.NewTypedToolHandler(s.ListSnippets),
	)
	s.server.AddTool(
		mcp.NewTool(
			constant.DeleteSnippet,
			mcp.WithDescription("Delete a snippet file from the Proxmox host"),
			mcp.WithString(
				"name",
				mcp.Required(),
				mcp.Description("Snippet filename"),
			),
		),
		mcp.NewTypedToolHandler(s.DeleteSnippet),
	)
}
