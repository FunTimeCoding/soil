package constant

import (
	"github.com/funtimecoding/soil/pkg/identity"
	"time"
)

const (
	PollInterval       = time.Minute
	ActivePollInterval = 10 * time.Second
)
const (
	BoardTitle = "Pipelines"
	BoardPath  = "/"

	SuccessIcon = "/static/gitlab-success.png"
	FailIcon    = "/static/gitlab-fail.png"
	RunningIcon = "/static/gitlab-running.png"
	PendingIcon = "/static/gitlab-pending.png"
	WarningIcon = "/static/gitlab-warning.png"

	RenovatePrefix = "renovate/"

	BoardEvent   = "board"
	SummaryEvent = "summary"
)

var Identity = identity.New(
	"gogitlabd",
	"GitLab API bridge",
	"gogitlabd",
).WithInstructions(
	"GitLab API - projects, pipelines, merge requests, commits, CI variables. Projects are referenced by path (owner/repo) or numeric ID.",
)

const (
	GetProject         = "get_project"
	ListProjects       = "list_projects"
	SearchRepositories = "search_repositories"
	GetRepositoryTree  = "get_repository_tree"
	GetFileContents    = "get_file_contents"

	ListPipelines        = "list_pipelines"
	GetPipeline          = "get_pipeline"
	CreatePipeline       = "create_pipeline"
	ListPipelineJobs     = "list_pipeline_jobs"
	GetPipelineJob       = "get_pipeline_job"
	GetPipelineJobOutput = "get_pipeline_job_output"
	RetryPipeline        = "retry_pipeline"
	RetryPipelineJob     = "retry_pipeline_job"
	CancelPipeline       = "cancel_pipeline"
	CancelPipelineJob    = "cancel_pipeline_job"

	ListRegistryRepositories = "list_registry_repositories"
	DeleteRegistryRepository = "delete_registry_repository"

	ListMergeRequests       = "list_merge_requests"
	GetMergeRequest         = "get_merge_request"
	GetMergeRequestDiffs    = "get_merge_request_diffs"
	MergeRequestDiscussions = "mr_discussions"
	CreateMergeRequestNote  = "create_merge_request_note"

	ListCommits   = "list_commits"
	GetCommit     = "get_commit"
	GetCommitDiff = "get_commit_diff"

	ListProjectVariables  = "list_project_variables"
	GetProjectVariable    = "get_project_variable"
	CreateProjectVariable = "create_project_variable"
	UpdateProjectVariable = "update_project_variable"
	DeleteProjectVariable = "delete_project_variable"

	CreateBranch = "create_branch"
)
