package mark

import (
	"context"
	"fmt"
	library "github.com/funtimecoding/soil/pkg/constant"
	"github.com/funtimecoding/soil/pkg/generative/model_context/constant"
	"github.com/funtimecoding/soil/pkg/generative/model_context/example/mark/option"
	"github.com/funtimecoding/soil/pkg/generative/model_context/server"
	"github.com/funtimecoding/soil/pkg/project"
	"github.com/funtimecoding/soil/pkg/system"
	"github.com/funtimecoding/soil/pkg/web"
	webConstant "github.com/funtimecoding/soil/pkg/web/constant"
	"github.com/mark3labs/mcp-go/mcp"
	mark "github.com/mark3labs/mcp-go/server"
	"net/http"
	"os"
)

// Run Test payloads for curl/stdin (JSON-RPC over stdio or HTTP):
//
//	{"jsonrpc": "2.0", "id": 3, "method": "resources/list", "params": {}}
//	{"jsonrpc": "2.0", "id": 6, "method": "resources/read", "params": {"uri": "docs://readme"}}
//	{"jsonrpc": "2.0", "id": 6, "method": "resources/read", "params": {"uri": "users://123/profile"}}
//	{"jsonrpc": "2.0", "id": 6, "method": "resources/read", "params": {"uri": "users://122/profile"}}
func Run(o *option.Mark) {
	s := mark.NewMCPServer(
		"Demo",
		library.DefaultVersion,
		mark.WithToolCapabilities(false),
	)
	s.AddTool(
		mcp.NewTool(
			constant.GreetTool,
			mcp.WithDescription("Say hello to someone"),
			mcp.WithString(
				constant.NameParameter,
				mcp.Required(),
				mcp.Description("Name of the person to greet"),
			),
		),
		func(
			_ context.Context,
			r mcp.CallToolRequest,
		) (*mcp.CallToolResult, error) {
			name, e := r.RequireString(constant.NameParameter)

			if e != nil {
				return mcp.NewToolResultError(e.Error()), nil
			}

			return mcp.NewToolResultText(
				fmt.Sprintf("Hello %s", name),
			), nil
		},
	)
	s.AddResource(
		mcp.NewResource(
			ReadmeDocument,
			"Project README",
			mcp.WithResourceDescription(
				"The project's README file",
			),
			mcp.WithMIMEType(webConstant.Markdown),
		),
		func(
			_ context.Context,
			_ mcp.ReadResourceRequest,
		) ([]mcp.ResourceContents, error) {
			content, e := os.ReadFile(project.ReadmeFile)

			if e != nil {
				return nil, e
			}

			return []mcp.ResourceContents{
				mcp.TextResourceContents{
					URI:      ReadmeDocument,
					MIMEType: webConstant.Markdown,
					Text:     string(content),
				},
			}, nil
		},
	)
	s.AddResourceTemplate(
		mcp.NewResourceTemplate(
			"user://{id}/profile",
			"User Profile",
			mcp.WithTemplateDescription(
				"Returns user profile information",
			),
			mcp.WithTemplateMIMEType(webConstant.Object),
		),
		func(
			_ context.Context,
			r mcp.ReadResourceRequest,
		) ([]mcp.ResourceContents, error) {
			p, e := profile(parseLocator(r.Params.URI))

			if e != nil {
				return nil, e
			}

			return []mcp.ResourceContents{
				mcp.TextResourceContents{
					URI:      r.Params.URI,
					MIMEType: webConstant.Object,
					Text:     p,
				},
			}, nil
		},
	)

	if o.Local {
		server.New(s).ServeLocal()
		system.KillSignalBlock()
	} else {
		v := server.New(s)
		m := http.NewServeMux()
		v.Setup(m)
		h := web.Server(m, server.Address)
		web.ServeAsynchronous(h)
		system.KillSignalBlock()
		web.GracefulShutdown(context.Background(), h, true)
	}
}
