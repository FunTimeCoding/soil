package model_context

import (
	generativeConstant "github.com/funtimecoding/soil/pkg/generative/constant"
	"github.com/funtimecoding/soil/pkg/tool/gopnsensed/constant"
	"github.com/mark3labs/mcp-go/mcp"
)

func (s *Server) registerWrite() {
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
				generativeConstant.ParameterIdentifier,
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
				generativeConstant.ParameterIdentifier,
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
