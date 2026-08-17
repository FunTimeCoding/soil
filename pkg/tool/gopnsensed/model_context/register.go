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
}
