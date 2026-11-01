package model_context

import (
	argumentConstant "github.com/funtimecoding/soil/pkg/argument/constant"
	"github.com/funtimecoding/soil/pkg/tool/gogitlabd/constant"
	"github.com/mark3labs/mcp-go/mcp"
)

func (s *Server) registerPipelines() {
	s.server.AddTool(
		mcp.NewTool(
			constant.ListPipelines,
			mcp.WithDescription("List pipelines for a GitLab project"),
			mcp.WithString(
				argumentConstant.Project,
				mcp.Required(),
				mcp.Description("Project path (owner/repo) or ID"),
			),
			mcp.WithString(
				constant.ReferenceParameter,
				mcp.Description("Filter by branch or tag"),
			),
			mcp.WithString(
				"status",
				mcp.Description(
					"Filter by status (running, pending, success, failed, canceled)",
				),
			),
		),
		mcp.NewTypedToolHandler(s.ListPipelines),
	)
	s.server.AddTool(
		mcp.NewTool(
			constant.GetPipeline,
			mcp.WithDescription("Get details of a GitLab pipeline"),
			mcp.WithString(
				argumentConstant.Project,
				mcp.Required(),
				mcp.Description("Project path (owner/repo) or ID"),
			),
			mcp.WithNumber(
				argumentConstant.Pipeline,
				mcp.Required(),
				mcp.Description("Pipeline ID"),
			),
		),
		mcp.NewTypedToolHandler(s.GetPipeline),
	)
	s.server.AddTool(
		mcp.NewTool(
			constant.CreatePipeline,
			mcp.WithDescription("Trigger a new pipeline for a GitLab project"),
			mcp.WithString(
				argumentConstant.Project,
				mcp.Required(),
				mcp.Description("Project path (owner/repo) or ID"),
			),
			mcp.WithString(
				constant.ReferenceParameter,
				mcp.Required(),
				mcp.Description("Branch or tag to run the pipeline on"),
			),
			mcp.WithObject(
				"variables",
				mcp.Description(
					"Pipeline variables as key-value pairs (e.g. {\"DEPLOY_TARGET\": \"sentry\"})",
				),
			),
		),
		mcp.NewTypedToolHandler(s.CreatePipeline),
	)
	s.server.AddTool(
		mcp.NewTool(
			constant.RetryPipeline,
			mcp.WithDescription("Retry a failed GitLab pipeline"),
			mcp.WithString(
				argumentConstant.Project,
				mcp.Required(),
				mcp.Description("Project path (owner/repo) or ID"),
			),
			mcp.WithNumber(
				argumentConstant.Pipeline,
				mcp.Required(),
				mcp.Description("Pipeline ID"),
			),
		),
		mcp.NewTypedToolHandler(s.RetryPipeline),
	)
	s.server.AddTool(
		mcp.NewTool(
			constant.CancelPipeline,
			mcp.WithDescription("Cancel a running GitLab pipeline"),
			mcp.WithString(
				argumentConstant.Project,
				mcp.Required(),
				mcp.Description("Project path (owner/repo) or ID"),
			),
			mcp.WithNumber(
				argumentConstant.Pipeline,
				mcp.Required(),
				mcp.Description("Pipeline ID"),
			),
		),
		mcp.NewTypedToolHandler(s.CancelPipeline),
	)
}
