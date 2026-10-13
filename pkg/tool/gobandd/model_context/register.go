package model_context

import (
	"github.com/funtimecoding/soil/pkg/tool/gobandd/constant"
	"github.com/mark3labs/mcp-go/mcp"
)

func (s *Server) register() {
	s.instances.Register(s.server)
	s.server.AddTool(
		mcp.NewTool(
			constant.Status,
			mcp.WithDescription(
				"Read the active engine's provisioning, redirection, consent and power state, and the power operations the firmware accepts now.",
			),
		),
		mcp.NewTypedToolHandler(s.status),
	)
	s.server.AddTool(
		mcp.NewTool(
			constant.Start,
			mcp.WithDescription(
				"Power on a machine that is off. Refused while the machine is on.",
			),
		),
		mcp.NewTypedToolHandler(s.start),
	)
	s.server.AddTool(
		mcp.NewTool(
			constant.Shutdown,
			mcp.WithDescription(
				"Ask the operating system to shut down through LMS - its services and guests stop cleanly. Refused when the firmware does not offer the graceful state.",
			),
			mcp.WithBoolean(
				constant.ForceParameter,
				mcp.Description(
					"Cut power below the operating system instead - nothing gets to stop, unwritten data is lost. Only when a graceful shutdown was accepted and the machine never went down.",
				),
			),
		),
		mcp.NewTypedToolHandler(s.shutdown),
	)
	s.server.AddTool(
		mcp.NewTool(
			constant.Reboot,
			mcp.WithDescription(
				"Ask the operating system to restart through LMS - its services and guests stop cleanly. Refused when the firmware does not offer the graceful state.",
			),
			mcp.WithBoolean(
				constant.ForceParameter,
				mcp.Description(
					"Power cycle below the operating system instead - nothing gets to stop, unwritten data is lost. Only when a graceful restart was accepted and the machine never went down.",
				),
			),
		),
		mcp.NewTypedToolHandler(s.reboot),
	)
	s.server.AddTool(
		mcp.NewTool(
			constant.EnableRedirection,
			mcp.WithDescription(
				"Switch on the active engine's redirection listener and screen service where they are off, so an AMT-aware client can open screen and serial sessions. Writes only what is missing.",
			),
		),
		mcp.NewTypedToolHandler(s.enableRedirection),
	)
}
