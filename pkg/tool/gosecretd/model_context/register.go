package model_context

import (
	generativeConstant "github.com/funtimecoding/soil/pkg/generative/constant"
	"github.com/funtimecoding/soil/pkg/tool/gosecretd/constant"
	"github.com/mark3labs/mcp-go/mcp"
)

func (s *Server) register() {
	s.server.AddTool(
		mcp.NewTool(
			constant.Status,
			mcp.WithDescription(
				"Whether the secret store is initialized and sealed, and its version. A sealed store serves nothing until unsealed.",
			),
		),
		mcp.NewTypedToolHandler(s.Status),
	)
	s.server.AddTool(
		mcp.NewTool(
			constant.ListSecrets,
			mcp.WithDescription(
				"List secrets and folders below a path; folders end in a slash.",
			),
			mcp.WithString(
				constant.PathParameter,
				mcp.Description("Folder to list, empty for the top level."),
			),
		),
		mcp.NewTypedToolHandler(s.ListSecrets),
	)
	s.server.AddTool(
		mcp.NewTool(
			constant.ShowSecret,
			mcp.WithDescription(
				"Show a secret's key names. Values stay masked; reveal one with reveal_secret.",
			),
			mcp.WithString(
				constant.PathParameter,
				mcp.Required(),
				mcp.Description("Secret path: <namespace>/<secret-name>."),
			),
		),
		mcp.NewTypedToolHandler(s.ShowSecret),
	)
	s.server.AddTool(
		mcp.NewTool(
			constant.RevealSecret,
			mcp.WithDescription(
				"Reveal one key's value. Call only when the value itself is needed.",
			),
			mcp.WithString(
				constant.PathParameter,
				mcp.Required(),
				mcp.Description("Secret path: <namespace>/<secret-name>."),
			),
			mcp.WithString(
				generativeConstant.ParameterKey,
				mcp.Required(),
				mcp.Description("The one key to reveal."),
			),
		),
		mcp.NewTypedToolHandler(s.RevealSecret),
	)
	s.server.AddTool(
		mcp.NewTool(
			constant.PutSecret,
			mcp.WithDescription(
				"Merge keys into a secret, creating it when absent. Existing keys not named here stay as they are.",
			),
			mcp.WithString(
				constant.PathParameter,
				mcp.Required(),
				mcp.Description("Secret path: <namespace>/<secret-name>."),
			),
			mcp.WithObject(
				constant.ValuesParameter,
				mcp.Required(),
				mcp.Description("Keys and their values."),
				mcp.AdditionalProperties(map[string]any{"type": "string"}),
			),
		),
		mcp.NewTypedToolHandler(s.PutSecret),
	)
}
