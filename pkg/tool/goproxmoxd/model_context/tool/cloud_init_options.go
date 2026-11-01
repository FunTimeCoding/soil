package tool

import "github.com/mark3labs/mcp-go/mcp"

func cloudInitOptions() []mcp.ToolOption {
	return []mcp.ToolOption{
		mcp.WithString("ci_user", mcp.Description("Cloud-init user")),
		mcp.WithString("ci_password", mcp.Description("Cloud-init password")),
		mcp.WithString(
			"ssh_keys",
			mcp.Description("SSH public keys, newline-separated"),
		),
		mcp.WithString(
			"ip_config",
			mcp.Description(
				"Cloud-init IP config (default ip=dhcp when cloud-init is used)",
			),
		),
		mcp.WithString(
			"search_domain",
			mcp.Description("DNS search domain for cloud-init, e.g. local"),
		),
	}
}
