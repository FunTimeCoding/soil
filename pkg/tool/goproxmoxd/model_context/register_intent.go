package model_context

import (
	"github.com/funtimecoding/soil/pkg/tool/goproxmoxd/constant"
	"github.com/mark3labs/mcp-go/mcp"
)

func (s *Server) registerIntent() {
	s.server.AddTool(
		mcp.NewTool(
			constant.Reconcile,
			mcp.WithDescription(
				"Compare the hypervisor's guests with the power intent netbox records: running guests recorded offline become active, and with full, stopped guests recorded active become offline. Guests netbox does not know are listed, never created. Nothing is written unless write is set.",
			),
			mcp.WithBoolean(
				constant.FullParameter,
				mcp.Description("Also record stopped guests as offline"),
			),
			mcp.WithBoolean(
				constant.WriteParameter,
				mcp.Description(
					"Write the changes; without it they are only listed",
				),
			),
		),
		mcp.NewTypedToolHandler(s.Reconcile),
	)
}
