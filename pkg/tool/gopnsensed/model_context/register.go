package model_context

import (
	generative "github.com/funtimecoding/soil/pkg/generative/constant"
	"github.com/funtimecoding/soil/pkg/tool/gopnsensed/constant"
	"github.com/mark3labs/mcp-go/mcp"
)

func (s *Server) register() {
	query := mcp.WithString(
		generative.ParameterQuery,
		mcp.Description("Search phrase to filter results."),
	)
	s.server.AddTool(
		mcp.NewTool(
			constant.ListLeases,
			mcp.WithDescription(
				"List DHCP leases with hostname, hardware address, and whether a static reservation covers them.",
			),
			query,
		),
		s.listLeases,
	)
	s.server.AddTool(
		mcp.NewTool(
			constant.ListHosts,
			mcp.WithDescription(
				"List Dnsmasq host entries - static DHCP reservations and manual name mappings.",
			),
			query,
		),
		s.listHosts,
	)
	s.server.AddTool(
		mcp.NewTool(
			constant.ListPools,
			mcp.WithDescription("List DHCP pools with their address ranges."),
			query,
		),
		s.listPools,
	)
	s.server.AddTool(
		mcp.NewTool(
			constant.ListRules,
			mcp.WithDescription(
				"List firewall filter rules with interface, direction, action, networks, ports, and log flag.",
			),
			query,
		),
		s.listRules,
	)
	s.server.AddTool(
		mcp.NewTool(
			constant.ListAliases,
			mcp.WithDescription("List firewall aliases and their content."),
			query,
		),
		s.listAliases,
	)
	s.server.AddTool(
		mcp.NewTool(
			constant.ListSourceNat,
			mcp.WithDescription("List source NAT rules."),
			query,
		),
		s.listSourceNat,
	)
	s.server.AddTool(
		mcp.NewTool(
			constant.ListForwards,
			mcp.WithDescription(
				"List Unbound query forwards - which domains resolve through which servers.",
			),
			query,
		),
		s.listForwards,
	)
	s.server.AddTool(
		mcp.NewTool(
			constant.ListBlocklists,
			mcp.WithDescription("List Unbound DNS blocklists."),
			query,
		),
		s.listBlocklists,
	)
	s.server.AddTool(
		mcp.NewTool(
			constant.ListInterfaces,
			mcp.WithDescription(
				"List network interfaces with status, media, MAC, and addresses.",
			),
		),
		s.listInterfaces,
	)
	s.server.AddTool(
		mcp.NewTool(
			constant.FirewallLog,
			mcp.WithDescription(
				"Read recent firewall log records - matched rule, action, source, and destination.",
			),
			mcp.WithNumber(
				generative.ParameterLimit,
				mcp.Description("Maximum number of records, default 100."),
			),
		),
		s.firewallLog,
	)
	s.server.AddTool(
		mcp.NewTool(
			constant.FirewallStates,
			mcp.WithDescription(
				"Query the firewall state table - live connections with rule, age, and traffic counters.",
			),
			query,
		),
		s.firewallStates,
	)
	apply := mcp.WithBoolean(
		constant.ParameterApply,
		mcp.Description(
			"Reconfigure Dnsmasq after the write so it takes effect, default true. Pass false to batch several writes and reconfigure once.",
		),
	)
	s.server.AddTool(
		mcp.NewTool(
			constant.AddHost,
			mcp.WithDescription(
				"Add a Dnsmasq host entry - a static DHCP reservation, a name mapping, or both. Returns the new entry identifier.",
			),
			mcp.WithString(
				constant.ParameterHost,
				mcp.Description("Hostname without the domain part."),
			),
			mcp.WithString(
				constant.ParameterDomain,
				mcp.Description("Domain of the host."),
			),
			mcp.WithString(
				constant.ParameterAddress,
				mcp.Description("Address to map or reserve."),
			),
			mcp.WithString(
				constant.ParameterHardwareAddress,
				mcp.Description("MAC address the reservation binds to."),
			),
			mcp.WithString(
				constant.ParameterClientIdentifier,
				mcp.Description("DHCP client identifier."),
			),
			mcp.WithString(
				constant.ParameterDescription,
				mcp.Description("Description of the entry."),
			),
			apply,
		),
		s.addHost,
	)
	s.server.AddTool(
		mcp.NewTool(
			constant.SetHost,
			mcp.WithDescription(
				"Update a Dnsmasq host entry. Omitted fields keep their current value; passing an empty one clears it.",
			),
			mcp.WithString(
				generative.ParameterIdentifier,
				mcp.Required(),
				mcp.Description("Identifier of the entry to update."),
			),
			mcp.WithString(
				constant.ParameterHost,
				mcp.Description("Hostname without the domain part."),
			),
			mcp.WithString(
				constant.ParameterDomain,
				mcp.Description("Domain of the host."),
			),
			mcp.WithString(
				constant.ParameterAddress,
				mcp.Description("Address to map or reserve."),
			),
			mcp.WithString(
				constant.ParameterHardwareAddress,
				mcp.Description("MAC address the reservation binds to."),
			),
			mcp.WithString(
				constant.ParameterClientIdentifier,
				mcp.Description("DHCP client identifier."),
			),
			mcp.WithString(
				constant.ParameterDescription,
				mcp.Description("Description of the entry."),
			),
			apply,
		),
		s.setHost,
	)
	s.server.AddTool(
		mcp.NewTool(
			constant.DeleteHost,
			mcp.WithDescription(
				"Delete one Dnsmasq host entry by identifier. Removal is permanent; the entry is not recoverable from here.",
			),
			mcp.WithString(
				generative.ParameterIdentifier,
				mcp.Required(),
				mcp.Description("Identifier of the entry to delete."),
			),
			apply,
		),
		s.deleteHost,
	)
	s.server.AddTool(
		mcp.NewTool(
			constant.ReconfigureDnsmasq,
			mcp.WithDescription(
				"Apply pending Dnsmasq configuration - required after deferred writes, and reloads the firewall filter.",
			),
		),
		s.reconfigureDnsmasq,
	)
}
