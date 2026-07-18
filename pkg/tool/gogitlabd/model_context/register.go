package model_context

import (
	"github.com/funtimecoding/soil/pkg/tool/gogitlabd/constant"
	"github.com/mark3labs/mcp-go/mcp"
)

func (s *Server) register() {
	s.server.AddTool(
		mcp.NewTool(
			constant.GetProject,
			mcp.WithDescription("Get a GitLab project by path or ID"),
			mcp.WithString(
				"project",
				mcp.Required(),
				mcp.Description("Project path (owner/repo) or ID"),
			),
		),
		mcp.NewTypedToolHandler(s.GetProject),
	)
	s.server.AddTool(
		mcp.NewTool(
			constant.ListProjects,
			mcp.WithDescription("List accessible GitLab projects"),
			mcp.WithString(
				"search",
				mcp.Description("Filter by name"),
			),
		),
		mcp.NewTypedToolHandler(s.ListProjects),
	)
	s.server.AddTool(
		mcp.NewTool(
			constant.SearchRepositories,
			mcp.WithDescription("Search GitLab projects by name"),
			mcp.WithString(
				"query",
				mcp.Required(),
				mcp.Description("Search query"),
			),
		),
		mcp.NewTypedToolHandler(s.SearchRepositories),
	)
	s.server.AddTool(
		mcp.NewTool(
			constant.GetRepositoryTree,
			mcp.WithDescription("List files and directories in a repository"),
			mcp.WithString(
				"project",
				mcp.Required(),
				mcp.Description("Project path (owner/repo) or ID"),
			),
			mcp.WithString(
				"path",
				mcp.Description("Directory path within the repository"),
			),
			mcp.WithString(
				"reference",
				mcp.Description("Branch name, tag, or commit SHA"),
			),
			mcp.WithBoolean(
				"recursive",
				mcp.Description("List files recursively"),
			),
		),
		mcp.NewTypedToolHandler(s.GetRepositoryTree),
	)
	s.server.AddTool(
		mcp.NewTool(
			constant.GetFileContents,
			mcp.WithDescription("Read a file from a GitLab repository"),
			mcp.WithString(
				"project",
				mcp.Required(),
				mcp.Description("Project path (owner/repo) or ID"),
			),
			mcp.WithString(
				"path",
				mcp.Required(),
				mcp.Description("File path within the repository"),
			),
			mcp.WithString(
				"reference",
				mcp.Description("Branch name, tag, or commit SHA"),
			),
		),
		mcp.NewTypedToolHandler(s.GetFileContents),
	)
	s.server.AddTool(
		mcp.NewTool(
			constant.ListPipelines,
			mcp.WithDescription("List pipelines for a GitLab project"),
			mcp.WithString(
				"project",
				mcp.Required(),
				mcp.Description("Project path (owner/repo) or ID"),
			),
			mcp.WithString(
				"reference",
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
				"project",
				mcp.Required(),
				mcp.Description("Project path (owner/repo) or ID"),
			),
			mcp.WithNumber(
				"pipeline",
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
				"project",
				mcp.Required(),
				mcp.Description("Project path (owner/repo) or ID"),
			),
			mcp.WithString(
				"reference",
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
			constant.ListPipelineJobs,
			mcp.WithDescription("List jobs in a GitLab pipeline"),
			mcp.WithString(
				"project",
				mcp.Required(),
				mcp.Description("Project path (owner/repo) or ID"),
			),
			mcp.WithNumber(
				"pipeline",
				mcp.Required(),
				mcp.Description("Pipeline ID"),
			),
		),
		mcp.NewTypedToolHandler(s.ListPipelineJobs),
	)
	s.server.AddTool(
		mcp.NewTool(
			constant.GetPipelineJob,
			mcp.WithDescription("Get details of a pipeline job"),
			mcp.WithString(
				"project",
				mcp.Required(),
				mcp.Description("Project path (owner/repo) or ID"),
			),
			mcp.WithNumber(
				"job",
				mcp.Required(),
				mcp.Description("Job ID"),
			),
		),
		mcp.NewTypedToolHandler(s.GetPipelineJob),
	)
	s.server.AddTool(
		mcp.NewTool(
			constant.GetPipelineJobOutput,
			mcp.WithDescription("Read the log output of a pipeline job"),
			mcp.WithString(
				"project",
				mcp.Required(),
				mcp.Description("Project path (owner/repo) or ID"),
			),
			mcp.WithNumber(
				"job",
				mcp.Required(),
				mcp.Description("Job ID"),
			),
		),
		mcp.NewTypedToolHandler(s.GetPipelineJobOutput),
	)
	s.server.AddTool(
		mcp.NewTool(
			constant.RetryPipeline,
			mcp.WithDescription("Retry a failed GitLab pipeline"),
			mcp.WithString(
				"project",
				mcp.Required(),
				mcp.Description("Project path (owner/repo) or ID"),
			),
			mcp.WithNumber(
				"pipeline",
				mcp.Required(),
				mcp.Description("Pipeline ID"),
			),
		),
		mcp.NewTypedToolHandler(s.RetryPipeline),
	)
	s.server.AddTool(
		mcp.NewTool(
			constant.RetryPipelineJob,
			mcp.WithDescription("Retry a specific pipeline job"),
			mcp.WithString(
				"project",
				mcp.Required(),
				mcp.Description("Project path (owner/repo) or ID"),
			),
			mcp.WithNumber(
				"job",
				mcp.Required(),
				mcp.Description("Job ID"),
			),
		),
		mcp.NewTypedToolHandler(s.RetryPipelineJob),
	)
	s.server.AddTool(
		mcp.NewTool(
			constant.CancelPipeline,
			mcp.WithDescription("Cancel a running GitLab pipeline"),
			mcp.WithString(
				"project",
				mcp.Required(),
				mcp.Description("Project path (owner/repo) or ID"),
			),
			mcp.WithNumber(
				"pipeline",
				mcp.Required(),
				mcp.Description("Pipeline ID"),
			),
		),
		mcp.NewTypedToolHandler(s.CancelPipeline),
	)
	s.server.AddTool(
		mcp.NewTool(
			constant.CancelPipelineJob,
			mcp.WithDescription("Cancel a specific pipeline job"),
			mcp.WithString(
				"project",
				mcp.Required(),
				mcp.Description("Project path (owner/repo) or ID"),
			),
			mcp.WithNumber(
				"job",
				mcp.Required(),
				mcp.Description("Job ID"),
			),
		),
		mcp.NewTypedToolHandler(s.CancelPipelineJob),
	)
	s.server.AddTool(
		mcp.NewTool(
			constant.ListMergeRequests,
			mcp.WithDescription("List merge requests for a GitLab project"),
			mcp.WithString(
				"project",
				mcp.Required(),
				mcp.Description("Project path (owner/repo) or ID"),
			),
			mcp.WithString(
				"state",
				mcp.Description(
					"Filter by state (opened, closed, merged, all)",
				),
			),
		),
		mcp.NewTypedToolHandler(s.ListMergeRequests),
	)
	s.server.AddTool(
		mcp.NewTool(
			constant.GetMergeRequest,
			mcp.WithDescription("Get details of a merge request"),
			mcp.WithString(
				"project",
				mcp.Required(),
				mcp.Description("Project path (owner/repo) or ID"),
			),
			mcp.WithNumber(
				"merge_request",
				mcp.Required(),
				mcp.Description("Merge request IID"),
			),
		),
		mcp.NewTypedToolHandler(s.GetMergeRequest),
	)
	s.server.AddTool(
		mcp.NewTool(
			constant.GetMergeRequestDiffs,
			mcp.WithDescription("Get the diff of a merge request"),
			mcp.WithString(
				"project",
				mcp.Required(),
				mcp.Description("Project path (owner/repo) or ID"),
			),
			mcp.WithNumber(
				"merge_request",
				mcp.Required(),
				mcp.Description("Merge request IID"),
			),
		),
		mcp.NewTypedToolHandler(s.GetMergeRequestDiffs),
	)
	s.server.AddTool(
		mcp.NewTool(
			constant.MergeRequestDiscussions,
			mcp.WithDescription(
				"List discussion threads on a merge request",
			),
			mcp.WithString(
				"project",
				mcp.Required(),
				mcp.Description("Project path (owner/repo) or ID"),
			),
			mcp.WithNumber(
				"merge_request",
				mcp.Required(),
				mcp.Description("Merge request IID"),
			),
		),
		mcp.NewTypedToolHandler(s.MergeRequestDiscussions),
	)
	s.server.AddTool(
		mcp.NewTool(
			constant.CreateMergeRequestNote,
			mcp.WithDescription("Add a comment to a merge request"),
			mcp.WithString(
				"project",
				mcp.Required(),
				mcp.Description("Project path (owner/repo) or ID"),
			),
			mcp.WithNumber(
				"merge_request",
				mcp.Required(),
				mcp.Description("Merge request IID"),
			),
			mcp.WithString(
				"body",
				mcp.Required(),
				mcp.Description("Comment body"),
			),
		),
		mcp.NewTypedToolHandler(s.CreateMergeRequestNote),
	)
	s.server.AddTool(
		mcp.NewTool(
			constant.ListCommits,
			mcp.WithDescription("List commits for a GitLab project"),
			mcp.WithString(
				"project",
				mcp.Required(),
				mcp.Description("Project path (owner/repo) or ID"),
			),
			mcp.WithString(
				"reference",
				mcp.Description("Branch or tag name"),
			),
		),
		mcp.NewTypedToolHandler(s.ListCommits),
	)
	s.server.AddTool(
		mcp.NewTool(
			constant.GetCommit,
			mcp.WithDescription("Get details of a commit"),
			mcp.WithString(
				"project",
				mcp.Required(),
				mcp.Description("Project path (owner/repo) or ID"),
			),
			mcp.WithString(
				"sha",
				mcp.Required(),
				mcp.Description("Commit SHA"),
			),
		),
		mcp.NewTypedToolHandler(s.GetCommit),
	)
	s.server.AddTool(
		mcp.NewTool(
			constant.GetCommitDiff,
			mcp.WithDescription("Get the diff of a commit"),
			mcp.WithString(
				"project",
				mcp.Required(),
				mcp.Description("Project path (owner/repo) or ID"),
			),
			mcp.WithString(
				"sha",
				mcp.Required(),
				mcp.Description("Commit SHA"),
			),
		),
		mcp.NewTypedToolHandler(s.GetCommitDiff),
	)
	s.server.AddTool(
		mcp.NewTool(
			constant.ListProjectVariables,
			mcp.WithDescription(
				"List CI/CD variables for a GitLab project",
			),
			mcp.WithString(
				"project",
				mcp.Required(),
				mcp.Description("Project path (owner/repo) or ID"),
			),
		),
		mcp.NewTypedToolHandler(s.ListProjectVariables),
	)
	s.server.AddTool(
		mcp.NewTool(
			constant.GetProjectVariable,
			mcp.WithDescription("Get a CI/CD variable by key"),
			mcp.WithString(
				"project",
				mcp.Required(),
				mcp.Description("Project path (owner/repo) or ID"),
			),
			mcp.WithString(
				"key",
				mcp.Required(),
				mcp.Description("Variable key"),
			),
		),
		mcp.NewTypedToolHandler(s.GetProjectVariable),
	)
	s.server.AddTool(
		mcp.NewTool(
			constant.CreateProjectVariable,
			mcp.WithDescription("Create a CI/CD variable"),
			mcp.WithString(
				"project",
				mcp.Required(),
				mcp.Description("Project path (owner/repo) or ID"),
			),
			mcp.WithString(
				"key",
				mcp.Required(),
				mcp.Description("Variable key"),
			),
			mcp.WithString(
				"value",
				mcp.Required(),
				mcp.Description("Variable value"),
			),
			mcp.WithBoolean(
				"protected",
				mcp.Description(
					"Only expose to protected branches/tags",
				),
			),
			mcp.WithBoolean(
				"masked",
				mcp.Description("Mask value in job logs"),
			),
			mcp.WithBoolean(
				"expand",
				mcp.Description(
					"Expand $VARIABLE references in the value; enable only when the value actually contains a reference to expand (default: raw)",
				),
			),
		),
		mcp.NewTypedToolHandler(s.CreateProjectVariable),
	)
	s.server.AddTool(
		mcp.NewTool(
			constant.UpdateProjectVariable,
			mcp.WithDescription("Update a CI/CD variable"),
			mcp.WithString(
				"project",
				mcp.Required(),
				mcp.Description("Project path (owner/repo) or ID"),
			),
			mcp.WithString(
				"key",
				mcp.Required(),
				mcp.Description("Variable key"),
			),
			mcp.WithString(
				"value",
				mcp.Required(),
				mcp.Description("New variable value"),
			),
			mcp.WithBoolean(
				"protected",
				mcp.Description(
					"Only expose to protected branches/tags",
				),
			),
			mcp.WithBoolean(
				"masked",
				mcp.Description("Mask value in job logs"),
			),
			mcp.WithBoolean(
				"expand",
				mcp.Description(
					"Expand $VARIABLE references in the value; enable only when the value actually contains a reference to expand (default: raw)",
				),
			),
		),
		mcp.NewTypedToolHandler(s.UpdateProjectVariable),
	)
	s.server.AddTool(
		mcp.NewTool(
			constant.DeleteProjectVariable,
			mcp.WithDescription("Delete a CI/CD variable"),
			mcp.WithString(
				"project",
				mcp.Required(),
				mcp.Description("Project path (owner/repo) or ID"),
			),
			mcp.WithString(
				"key",
				mcp.Required(),
				mcp.Description("Variable key"),
			),
		),
		mcp.NewTypedToolHandler(s.DeleteProjectVariable),
	)
	s.server.AddTool(
		mcp.NewTool(
			constant.CreateBranch,
			mcp.WithDescription("Create a branch in a GitLab project"),
			mcp.WithString(
				"project",
				mcp.Required(),
				mcp.Description("Project path (owner/repo) or ID"),
			),
			mcp.WithString(
				"branch",
				mcp.Required(),
				mcp.Description("New branch name"),
			),
			mcp.WithString(
				"reference",
				mcp.Required(),
				mcp.Description(
					"Source branch, tag, or commit SHA to branch from",
				),
			),
		),
		mcp.NewTypedToolHandler(s.CreateBranch),
	)
}
