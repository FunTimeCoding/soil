package model_context

import (
	"github.com/funtimecoding/soil/pkg/tool/gogated/constant"
	"github.com/mark3labs/mcp-go/mcp"
)

func (s *Server) register() {
	s.server.AddTool(
		mcp.NewTool(
			constant.ListClient,
			mcp.WithDescription(
				"List every registered OAuth client with its redirect locators, scopes, grant types and authentication method. Secrets are never returned.",
			),
		),
		mcp.NewTypedToolHandler(s.listClient),
	)
	s.server.AddTool(
		mcp.NewTool(
			constant.CreateClient,
			mcp.WithDescription(
				"Register a relying party with this fleet's shape - authorization_code and refresh_token grants, client_secret_post, openid and offline scopes. The secret is returned once. For RFC 7591 dynamic registration with protocol defaults, POST /register instead.",
			),
			mcp.WithString(
				"redirect_locator",
				mcp.Required(),
				mcp.Description(
					"Callback locator, normally https://<host>/callback; several separated by spaces",
				),
			),
			mcp.WithString(
				"scope",
				mcp.Description(
					"Space-separated scopes; defaults to 'openid offline', which is what relying parties request",
				),
			),
		),
		mcp.NewTypedToolHandler(s.createClient),
	)
	s.server.AddTool(
		mcp.NewTool(
			constant.UpdateClient,
			mcp.WithDescription(
				"Change a registered client. Omitted fields are left untouched. Use this to conform a client created through RFC 7591 registration, whose protocol defaults differ from this fleet's.",
			),
			mcp.WithString(
				"identifier",
				mcp.Required(),
				mcp.Description("Client identifier"),
			),
			mcp.WithString(
				"redirect_locator",
				mcp.Description(
					"Replacement callback locators, space separated",
				),
			),
			mcp.WithString("scope", mcp.Description("Replacement scopes")),
			mcp.WithString(
				"grant_type",
				mcp.Description(
					"Replacement grant types, normally 'authorization_code refresh_token'",
				),
			),
			mcp.WithString(
				"response_type",
				mcp.Description("Replacement response types, normally 'code'"),
			),
			mcp.WithString(
				"token_endpoint_auth_method",
				mcp.Description(
					"Normally client_secret_post - the soil authorization client sends credentials in the form body",
				),
			),
		),
		mcp.NewTypedToolHandler(s.updateClient),
	)
	s.server.AddTool(
		mcp.NewTool(
			constant.DeleteClient,
			mcp.WithDescription(
				"Remove a registered client permanently. The relying party can no longer sign anyone in.",
			),
			mcp.WithString(
				"identifier",
				mcp.Required(),
				mcp.Description("Client identifier"),
			),
		),
		mcp.NewTypedToolHandler(s.deleteClient),
	)
}
