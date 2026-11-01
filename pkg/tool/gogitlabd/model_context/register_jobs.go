package model_context

import (
	argumentConstant "github.com/funtimecoding/soil/pkg/argument/constant"
	"github.com/funtimecoding/soil/pkg/generative/model_context/option"
	"github.com/funtimecoding/soil/pkg/tool/gogitlabd/constant"
	"github.com/mark3labs/mcp-go/mcp"
)

func (s *Server) registerJobs() {
	s.server.AddTool(
		mcp.NewTool(
			constant.ListPipelineJobs,
			mcp.WithDescription("List jobs in a GitLab pipeline"),
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
			option.Raw(),
		),
		mcp.NewTypedToolHandler(s.ListPipelineJobs),
	)
	s.server.AddTool(
		mcp.NewTool(
			constant.GetPipelineJob,
			mcp.WithDescription("Get details of a pipeline job"),
			mcp.WithString(
				argumentConstant.Project,
				mcp.Required(),
				mcp.Description("Project path (owner/repo) or ID"),
			),
			mcp.WithNumber(
				argumentConstant.Job,
				mcp.Required(),
				mcp.Description("Job ID"),
			),
			option.Raw(),
		),
		mcp.NewTypedToolHandler(s.GetPipelineJob),
	)
	s.server.AddTool(
		mcp.NewTool(
			constant.GetPipelineJobOutput,
			mcp.WithDescription("Read the log output of a pipeline job"),
			mcp.WithString(
				argumentConstant.Project,
				mcp.Required(),
				mcp.Description("Project path (owner/repo) or ID"),
			),
			mcp.WithNumber(
				argumentConstant.Job,
				mcp.Required(),
				mcp.Description("Job ID"),
			),
		),
		mcp.NewTypedToolHandler(s.GetPipelineJobOutput),
	)
	s.server.AddTool(
		mcp.NewTool(
			constant.RetryPipelineJob,
			mcp.WithDescription("Retry a specific pipeline job"),
			mcp.WithString(
				argumentConstant.Project,
				mcp.Required(),
				mcp.Description("Project path (owner/repo) or ID"),
			),
			mcp.WithNumber(
				argumentConstant.Job,
				mcp.Required(),
				mcp.Description("Job ID"),
			),
			option.Raw(),
		),
		mcp.NewTypedToolHandler(s.RetryPipelineJob),
	)
	s.server.AddTool(
		mcp.NewTool(
			constant.CancelPipelineJob,
			mcp.WithDescription("Cancel a specific pipeline job"),
			mcp.WithString(
				argumentConstant.Project,
				mcp.Required(),
				mcp.Description("Project path (owner/repo) or ID"),
			),
			mcp.WithNumber(
				argumentConstant.Job,
				mcp.Required(),
				mcp.Description("Job ID"),
			),
			option.Raw(),
		),
		mcp.NewTypedToolHandler(s.CancelPipelineJob),
	)
}
