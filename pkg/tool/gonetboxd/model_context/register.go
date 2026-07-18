package model_context

import (
	"github.com/funtimecoding/soil/pkg/generative/model_context/parameter"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/constant"
	"github.com/mark3labs/mcp-go/mcp"
)

func (s *Server) register() {
	s.server.AddTool(
		mcp.NewTool(
			constant.ListDevices,
			mcp.WithDescription(
				"List NetBox devices, optionally filtered by name",
			),
			mcp.WithString(
				parameter.Query,
				mcp.Description(
					"Filter by name (case-insensitive contains)",
				),
			),
		),
		s.listDevices,
	)
	s.server.AddTool(
		mcp.NewTool(
			constant.GetDevice,
			mcp.WithDescription(
				"Get a NetBox device by exact name",
			),
			mcp.WithString(
				parameter.Name,
				mcp.Required(),
				mcp.Description("Device name"),
			),
		),
		s.getDevice,
	)
	s.server.AddTool(
		mcp.NewTool(
			constant.ListInterfaces,
			mcp.WithDescription(
				"List interfaces for a NetBox device",
			),
			mcp.WithString(
				parameter.Name,
				mcp.Required(),
				mcp.Description("Device name"),
			),
		),
		s.listInterfaces,
	)
	s.server.AddTool(
		mcp.NewTool(
			constant.ListAddresses,
			mcp.WithDescription(
				"List IP addresses assigned to a NetBox device",
			),
			mcp.WithString(
				parameter.Name,
				mcp.Required(),
				mcp.Description("Device name"),
			),
		),
		s.listAddresses,
	)
	s.server.AddTool(
		mcp.NewTool(
			constant.ListSites,
			mcp.WithDescription("List all NetBox sites"),
		),
		s.listSites,
	)
	s.server.AddTool(
		mcp.NewTool(
			constant.ListTenants,
			mcp.WithDescription("List all NetBox tenants"),
		),
		s.listTenants,
	)
	s.server.AddTool(
		mcp.NewTool(
			constant.ListDeviceRoles,
			mcp.WithDescription("List all NetBox device roles"),
		),
		s.listDeviceRoles,
	)
	s.server.AddTool(
		mcp.NewTool(
			constant.ListManufacturers,
			mcp.WithDescription("List all NetBox manufacturers"),
		),
		s.listManufacturers,
	)
	s.server.AddTool(
		mcp.NewTool(
			constant.ListDeviceTypes,
			mcp.WithDescription("List all NetBox device types"),
		),
		s.listDeviceTypes,
	)
	s.server.AddTool(
		mcp.NewTool(
			constant.CreateSite,
			mcp.WithDescription("Create a NetBox site"),
			mcp.WithString(
				parameter.Name,
				mcp.Required(),
				mcp.Description("Site name"),
			),
		),
		s.createSite,
	)
	s.server.AddTool(
		mcp.NewTool(
			constant.CreateTenant,
			mcp.WithDescription("Create a NetBox tenant"),
			mcp.WithString(
				parameter.Name,
				mcp.Required(),
				mcp.Description("Tenant name"),
			),
		),
		s.createTenant,
	)
	s.server.AddTool(
		mcp.NewTool(
			constant.CreateDeviceRole,
			mcp.WithDescription("Create a NetBox device role"),
			mcp.WithString(
				parameter.Name,
				mcp.Required(),
				mcp.Description("Role name"),
			),
		),
		s.createDeviceRole,
	)
	s.server.AddTool(
		mcp.NewTool(
			constant.CreateManufacturer,
			mcp.WithDescription("Create a NetBox manufacturer"),
			mcp.WithString(
				parameter.Name,
				mcp.Required(),
				mcp.Description("Manufacturer name"),
			),
		),
		s.createManufacturer,
	)
	s.server.AddTool(
		mcp.NewTool(
			constant.CreateDeviceType,
			mcp.WithDescription("Create a NetBox device type"),
			mcp.WithString(
				constant.Model,
				mcp.Required(),
				mcp.Description("Model name"),
			),
			mcp.WithString(
				constant.Manufacturer,
				mcp.Required(),
				mcp.Description("Manufacturer name (must exist)"),
			),
		),
		s.createDeviceType,
	)
	s.server.AddTool(
		mcp.NewTool(
			constant.CreateDevice,
			mcp.WithDescription(
				"Create a NetBox device. Requires site, role, and device type to exist.",
			),
			mcp.WithString(
				parameter.Name,
				mcp.Required(),
				mcp.Description("Device name"),
			),
			mcp.WithString(
				constant.Role,
				mcp.Required(),
				mcp.Description("Device role name"),
			),
			mcp.WithString(
				constant.Type,
				mcp.Required(),
				mcp.Description("Device type model name"),
			),
			mcp.WithString(
				constant.Site,
				mcp.Required(),
				mcp.Description("Site name"),
			),
			mcp.WithString(
				constant.Tenant,
				mcp.Description("Tenant name (optional)"),
			),
		),
		s.createDevice,
	)
	s.server.AddTool(
		mcp.NewTool(
			constant.CreateInterface,
			mcp.WithDescription(
				"Create a network interface on a NetBox device",
			),
			mcp.WithString(
				constant.Device,
				mcp.Required(),
				mcp.Description("Device name"),
			),
			mcp.WithString(
				parameter.Name,
				mcp.Required(),
				mcp.Description("Interface name (e.g. eno1, eth0)"),
			),
			mcp.WithString(
				constant.Type,
				mcp.Required(),
				mcp.Description(
					"Interface type (e.g. 1000base-t, 10gbase-t)",
				),
			),
		),
		s.createInterface,
	)
	s.server.AddTool(
		mcp.NewTool(
			constant.CreateAddress,
			mcp.WithDescription(
				"Assign an IP address to a device interface",
			),
			mcp.WithString(
				constant.Device,
				mcp.Required(),
				mcp.Description("Device name"),
			),
			mcp.WithString(
				constant.Interface,
				mcp.Required(),
				mcp.Description("Interface name (e.g. eno1)"),
			),
			mcp.WithString(
				constant.Address,
				mcp.Required(),
				mcp.Description(
					"IP address in CIDR notation (e.g. 144.76.220.60/32)",
				),
			),
		),
		s.createAddress,
	)
	s.server.AddTool(
		mcp.NewTool(
			constant.ListClusterTypes,
			mcp.WithDescription("List all NetBox cluster types"),
		),
		s.listClusterTypes,
	)
	s.server.AddTool(
		mcp.NewTool(
			constant.ListClusters,
			mcp.WithDescription("List all NetBox clusters"),
		),
		s.listClusters,
	)
	s.server.AddTool(
		mcp.NewTool(
			constant.ListVirtualMachines,
			mcp.WithDescription(
				"List all NetBox virtual machines",
			),
		),
		s.listVirtualMachines,
	)
	s.server.AddTool(
		mcp.NewTool(
			constant.CreateClusterType,
			mcp.WithDescription("Create a NetBox cluster type"),
			mcp.WithString(
				parameter.Name,
				mcp.Required(),
				mcp.Description("Cluster type name"),
			),
		),
		s.createClusterType,
	)
	s.server.AddTool(
		mcp.NewTool(
			constant.CreateCluster,
			mcp.WithDescription(
				"Create a NetBox cluster (VM host grouping)",
			),
			mcp.WithString(
				parameter.Name,
				mcp.Required(),
				mcp.Description("Cluster name"),
			),
			mcp.WithString(
				constant.Type,
				mcp.Required(),
				mcp.Description("Cluster type name"),
			),
			mcp.WithString(
				constant.Site,
				mcp.Required(),
				mcp.Description("Site name"),
			),
		),
		s.createCluster,
	)
	s.server.AddTool(
		mcp.NewTool(
			constant.CreateVirtualMachine,
			mcp.WithDescription(
				"Create a NetBox virtual machine assigned to a cluster",
			),
			mcp.WithString(
				parameter.Name,
				mcp.Required(),
				mcp.Description("VM name"),
			),
			mcp.WithString(
				constant.Cluster,
				mcp.Required(),
				mcp.Description("Cluster name"),
			),
		),
		s.createVirtualMachine,
	)
	s.server.AddTool(
		mcp.NewTool(
			constant.CreateVirtualInterface,
			mcp.WithDescription(
				"Create an interface on a NetBox virtual machine",
			),
			mcp.WithString(
				constant.VirtualMachine,
				mcp.Required(),
				mcp.Description("VM name"),
			),
			mcp.WithString(
				parameter.Name,
				mcp.Required(),
				mcp.Description("Interface name (e.g. ens3)"),
			),
		),
		s.createVirtualInterface,
	)
	s.server.AddTool(
		mcp.NewTool(
			constant.CreateVirtualAddress,
			mcp.WithDescription(
				"Assign an IP address to a virtual machine interface",
			),
			mcp.WithString(
				constant.VirtualMachine,
				mcp.Required(),
				mcp.Description("VM name"),
			),
			mcp.WithString(
				constant.Interface,
				mcp.Required(),
				mcp.Description("Interface name"),
			),
			mcp.WithString(
				constant.Address,
				mcp.Required(),
				mcp.Description("IP in CIDR notation"),
			),
		),
		s.createVirtualAddress,
	)
	s.server.AddTool(
		mcp.NewTool(
			constant.ListTags,
			mcp.WithDescription("List all NetBox tags"),
		),
		s.listTags,
	)
	s.server.AddTool(
		mcp.NewTool(
			constant.CreateTag,
			mcp.WithDescription("Create a NetBox tag"),
			mcp.WithString(
				parameter.Name,
				mcp.Required(),
				mcp.Description("Tag name"),
			),
		),
		s.createTag,
	)
	s.server.AddTool(
		mcp.NewTool(
			constant.AddDeviceTag,
			mcp.WithDescription("Add a tag to a device"),
			mcp.WithString(
				constant.Device,
				mcp.Required(),
				mcp.Description("Device name"),
			),
			mcp.WithString(
				constant.Tag,
				mcp.Required(),
				mcp.Description("Tag name"),
			),
		),
		s.addDeviceTag,
	)
	s.server.AddTool(
		mcp.NewTool(
			constant.RemoveDeviceTag,
			mcp.WithDescription("Remove a tag from a device"),
			mcp.WithString(
				constant.Device,
				mcp.Required(),
				mcp.Description("Device name"),
			),
			mcp.WithString(
				constant.Tag,
				mcp.Required(),
				mcp.Description("Tag name"),
			),
		),
		s.removeDeviceTag,
	)
	s.server.AddTool(
		mcp.NewTool(
			constant.AddVirtualTag,
			mcp.WithDescription("Add a tag to a virtual machine"),
			mcp.WithString(
				constant.VirtualMachine,
				mcp.Required(),
				mcp.Description("VM name"),
			),
			mcp.WithString(
				constant.Tag,
				mcp.Required(),
				mcp.Description("Tag name"),
			),
		),
		s.addVirtualTag,
	)
	s.server.AddTool(
		mcp.NewTool(
			constant.RemoveVirtualTag,
			mcp.WithDescription(
				"Remove a tag from a virtual machine",
			),
			mcp.WithString(
				constant.VirtualMachine,
				mcp.Required(),
				mcp.Description("VM name"),
			),
			mcp.WithString(
				constant.Tag,
				mcp.Required(),
				mcp.Description("Tag name"),
			),
		),
		s.removeVirtualTag,
	)
	s.server.AddTool(
		mcp.NewTool(
			constant.AddDeviceJournalEntry,
			mcp.WithDescription(
				"Add a journal entry to a device - an append-only timestamped note, visible on the device page in NetBox",
			),
			mcp.WithString(
				constant.Device,
				mcp.Required(),
				mcp.Description("Device name"),
			),
			mcp.WithString(
				constant.Comments,
				mcp.Required(),
				mcp.Description("Entry text, markdown"),
			),
			mcp.WithString(
				constant.Kind,
				mcp.Description(
					"Entry kind: info, success, warning, or danger",
				),
			),
		),
		s.addDeviceJournalEntry,
	)
	s.server.AddTool(
		mcp.NewTool(
			constant.ListDeviceJournalEntries,
			mcp.WithDescription(
				"List journal entries for a device, newest first",
			),
			mcp.WithString(
				constant.Device,
				mcp.Required(),
				mcp.Description("Device name"),
			),
			mcp.WithNumber(
				parameter.Limit,
				mcp.Description("Maximum entries to return"),
			),
			mcp.WithNumber(
				parameter.Offset,
				mcp.Description("Entries to skip"),
			),
		),
		s.listDeviceJournalEntries,
	)
	s.server.AddTool(
		mcp.NewTool(
			constant.AddVirtualJournalEntry,
			mcp.WithDescription(
				"Add a journal entry to a virtual machine - an append-only timestamped note, visible on the VM page in NetBox",
			),
			mcp.WithString(
				constant.VirtualMachine,
				mcp.Required(),
				mcp.Description("VM name"),
			),
			mcp.WithString(
				constant.Comments,
				mcp.Required(),
				mcp.Description("Entry text, markdown"),
			),
			mcp.WithString(
				constant.Kind,
				mcp.Description(
					"Entry kind: info, success, warning, or danger",
				),
			),
		),
		s.addVirtualJournalEntry,
	)
	s.server.AddTool(
		mcp.NewTool(
			constant.ListVirtualJournalEntries,
			mcp.WithDescription(
				"List journal entries for a virtual machine, newest first",
			),
			mcp.WithString(
				constant.VirtualMachine,
				mcp.Required(),
				mcp.Description("VM name"),
			),
			mcp.WithNumber(
				parameter.Limit,
				mcp.Description("Maximum entries to return"),
			),
			mcp.WithNumber(
				parameter.Offset,
				mcp.Description("Entries to skip"),
			),
		),
		s.listVirtualJournalEntries,
	)
	s.server.AddTool(
		mcp.NewTool(
			constant.UpdateJournalEntry,
			mcp.WithDescription(
				"Update a journal entry's comments or kind",
			),
			mcp.WithNumber(
				parameter.Identifier,
				mcp.Required(),
				mcp.Description("Journal entry identifier"),
			),
			mcp.WithString(
				constant.Comments,
				mcp.Description("New entry text, markdown"),
			),
			mcp.WithString(
				constant.Kind,
				mcp.Description(
					"New entry kind: info, success, warning, or danger",
				),
			),
		),
		s.updateJournalEntry,
	)
	s.server.AddTool(
		mcp.NewTool(
			constant.DeleteJournalEntry,
			mcp.WithDescription("Delete a journal entry"),
			mcp.WithNumber(
				parameter.Identifier,
				mcp.Required(),
				mcp.Description("Journal entry identifier"),
			),
		),
		s.deleteJournalEntry,
	)
	s.server.AddTool(
		mcp.NewTool(
			constant.ListTunnelGroups,
			mcp.WithDescription("List all VPN tunnel groups"),
		),
		s.listTunnelGroups,
	)
	s.server.AddTool(
		mcp.NewTool(
			constant.ListTunnels,
			mcp.WithDescription("List all VPN tunnels"),
		),
		s.listTunnels,
	)
	s.server.AddTool(
		mcp.NewTool(
			constant.ListTunnelTerminations,
			mcp.WithDescription("List all VPN tunnel terminations"),
		),
		s.listTunnelTerminations,
	)
	s.server.AddTool(
		mcp.NewTool(
			constant.CreateTunnelGroup,
			mcp.WithDescription("Create a VPN tunnel group"),
			mcp.WithString(
				parameter.Name,
				mcp.Required(),
				mcp.Description("Tunnel group name"),
			),
		),
		s.createTunnelGroup,
	)
	s.server.AddTool(
		mcp.NewTool(
			constant.CreateTunnel,
			mcp.WithDescription("Create a VPN tunnel in a group"),
			mcp.WithString(
				parameter.Name,
				mcp.Required(),
				mcp.Description("Tunnel name"),
			),
			mcp.WithString(
				constant.Encapsulation,
				mcp.Required(),
				mcp.Description(
					"Encapsulation type (e.g. wireguard, gre, openvpn)",
				),
			),
			mcp.WithString(
				constant.Group,
				mcp.Required(),
				mcp.Description("Tunnel group name"),
			),
		),
		s.createTunnel,
	)
	s.server.AddTool(
		mcp.NewTool(
			constant.CreateDeviceTunnelTermination,
			mcp.WithDescription(
				"Create a tunnel termination on a device interface",
			),
			mcp.WithString(
				constant.Device,
				mcp.Required(),
				mcp.Description("Device name"),
			),
			mcp.WithString(
				constant.Tunnel,
				mcp.Required(),
				mcp.Description("Tunnel name"),
			),
			mcp.WithString(
				constant.Interface,
				mcp.Required(),
				mcp.Description("Interface name on the device"),
			),
			mcp.WithString(
				constant.Role,
				mcp.Required(),
				mcp.Description("Termination role (peer, hub, spoke)"),
			),
		),
		s.createDeviceTunnelTermination,
	)
	s.server.AddTool(
		mcp.NewTool(
			constant.CreateVirtualTunnelTermination,
			mcp.WithDescription(
				"Create a tunnel termination on a VM interface",
			),
			mcp.WithString(
				constant.VirtualMachine,
				mcp.Required(),
				mcp.Description("VM name"),
			),
			mcp.WithString(
				constant.Tunnel,
				mcp.Required(),
				mcp.Description("Tunnel name"),
			),
			mcp.WithString(
				constant.Interface,
				mcp.Required(),
				mcp.Description("Interface name on the VM"),
			),
			mcp.WithString(
				constant.Role,
				mcp.Required(),
				mcp.Description("Termination role (peer, hub, spoke)"),
			),
		),
		s.createVirtualTunnelTermination,
	)
	s.server.AddTool(
		mcp.NewTool(
			constant.ListPrefixes,
			mcp.WithDescription("List all IP prefixes"),
		),
		s.listPrefixes,
	)
	s.server.AddTool(
		mcp.NewTool(
			constant.CreatePrefix,
			mcp.WithDescription("Create an IP prefix (subnet)"),
			mcp.WithString(
				constant.Prefix,
				mcp.Required(),
				mcp.Description(
					"CIDR notation (e.g. 192.168.178.0/24)",
				),
			),
			mcp.WithString(
				constant.Site,
				mcp.Description("Site name (optional)"),
			),
			mcp.WithString(
				constant.Description,
				mcp.Description("Description (optional)"),
			),
		),
		s.createPrefix,
	)
}
