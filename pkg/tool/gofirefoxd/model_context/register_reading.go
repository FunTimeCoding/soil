package model_context

import (
	"github.com/funtimecoding/soil/pkg/tool/gofirefoxd/constant"
	"github.com/mark3labs/mcp-go/mcp"
)

func (s *Server) registerReading() {
	s.server.AddTool(
		mcp.NewTool(
			constant.ReadTab,
			mcp.WithDescription(
				"Read the text content of a Firefox tab. Extracts readable article text by default (strips navigation, ads, footers). Falls back to raw innerText if readability extraction fails. Pass raw=true to skip readability and get innerText directly - better for dashboards, web apps, and non-article pages.",
			),
			mcp.WithNumber(
				"tab_id",
				mcp.Required(),
				mcp.Description("Firefox tab identifier"),
			),
			mcp.WithBoolean(
				"raw",
				mcp.Description(
					"Skip readability, return raw innerText (default false)",
				),
			),
		),
		mcp.NewTypedToolHandler(s.ReadTab),
	)
	s.server.AddTool(
		mcp.NewTool(
			constant.Screenshot,
			mcp.WithDescription(
				"Capture a PNG of a Firefox tab and write it to the download directory, returning the path to read. Use it to see what text extraction cannot: CAPTCHA and Cloudflare interstitials, which are frequent over Tor and look like an empty page otherwise. Firefox can only capture the visible tab, so this activates the target tab first.",
			),
			mcp.WithNumber(
				"tab_id",
				mcp.Required(),
				mcp.Description("Firefox tab identifier"),
			),
		),
		mcp.NewTypedToolHandler(s.Screenshot),
	)
}
