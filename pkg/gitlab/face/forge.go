package face

import (
	"github.com/funtimecoding/soil/pkg/gitlab/branch"
	"github.com/funtimecoding/soil/pkg/gitlab/commit"
	"github.com/funtimecoding/soil/pkg/gitlab/diff"
	"github.com/funtimecoding/soil/pkg/gitlab/discussion"
	"github.com/funtimecoding/soil/pkg/gitlab/file"
	"github.com/funtimecoding/soil/pkg/gitlab/image"
	"github.com/funtimecoding/soil/pkg/gitlab/job"
	"github.com/funtimecoding/soil/pkg/gitlab/merge_request"
	"github.com/funtimecoding/soil/pkg/gitlab/merge_request_detail"
	"github.com/funtimecoding/soil/pkg/gitlab/note"
	"github.com/funtimecoding/soil/pkg/gitlab/pipeline"
	"github.com/funtimecoding/soil/pkg/gitlab/pipeline_detail"
	"github.com/funtimecoding/soil/pkg/gitlab/project"
	"github.com/funtimecoding/soil/pkg/gitlab/registry_repository"
	"github.com/funtimecoding/soil/pkg/gitlab/tag"
	"github.com/funtimecoding/soil/pkg/gitlab/tree"
	"github.com/funtimecoding/soil/pkg/gitlab/variable"
	"gitlab.com/gitlab-org/api/client-go/v2"
)

type Forge interface {
	Project(identifier int64) (*project.Project, error)
	Projects() ([]*project.Project, error)
	MustProjects() []*project.Project
	Branches(project int64) ([]*branch.Branch, error)
	MustBranches(project int64) []*branch.Branch
	DeletePipeline(
		project int64,
		identifier int64,
	) error
	Tags(project int64) ([]*tag.Tag, error)
	MustTags(project int64) []*tag.Tag
	CreateTag(
		project int64,
		name string,
		reference string,
		message string,
	) (*tag.Tag, error)
	MustCreateTag(
		project int64,
		name string,
		reference string,
		message string,
	) *tag.Tag
	Commit(
		project int64,
		branch string,
		text string,
		path string,
		content string,
		update bool,
	) (*commit.Commit, error)
	MustCommit(
		project int64,
		branch string,
		text string,
		path string,
		content string,
		update bool,
	) *commit.Commit
	SearchProject(query string) ([]*project.Project, error)
	ProjectMergeRequests(
		project int64,
		state string,
	) ([]*merge_request.Request, error)
	Variables(project int64) ([]*variable.Variable, error)
	ProjectVariable(
		project int64,
		key string,
	) (*variable.Variable, error)
	CreateProjectVariable(
		project int64,
		key string,
		value string,
		protected bool,
		masked bool,
		literal bool,
	) (*variable.Variable, error)
	UpdateProjectVariable(
		project int64,
		key string,
		value string,
		protected bool,
		masked bool,
		literal bool,
	) (*variable.Variable, error)
	DeleteProjectVariable(
		project int64,
		key string,
	) error
	MergeRequest(
		project int64,
		identifier int64,
	) (*merge_request_detail.Detail, error)
	MergeRequestDiffs(
		project int64,
		identifier int64,
	) ([]*diff.Diff, error)
	MergeRequestDiscussions(
		project int64,
		identifier int64,
	) ([]*discussion.Discussion, error)
	CreateMergeRequestNote(
		project int64,
		identifier int64,
		body string,
	) (*note.Note, error)
	ReadCommit(
		project int64,
		sha string,
	) (*commit.Commit, error)
	ListCommits(
		project int64,
		reference string,
		limit int64,
	) ([]*commit.Commit, error)
	CommitDiff(
		project int64,
		sha string,
		limit int64,
	) ([]*diff.Diff, error)
	PipelineJob(
		project int64,
		identifier int64,
	) (*job.Job, error)
	PipelineJobs(
		project int64,
		identifier int64,
	) ([]*job.Job, error)
	CancelJob(
		project int64,
		identifier int64,
	) (*job.Job, error)
	Retry(
		project int64,
		jobIdentifier int64,
	) (*job.Job, error)
	Trace(
		project int64,
		jobIdentifier int64,
	) (string, error)
	CompareFiles(
		project int64,
		from string,
		to string,
	) ([]string, error)
	ResolveProject(identifier string) (int64, error)
	Pipeline(
		project int64,
		identifier int64,
	) (*pipeline_detail.Detail, error)
	Pipelines(
		project int64,
		reference string,
		status string,
		limit int64,
	) ([]*pipeline.Pipeline, error)
	MustPipelines(project int64) []*pipeline.Pipeline
	CreatePipeline(
		project int64,
		reference string,
		v []*gitlab.PipelineVariableOptions,
	) (*pipeline_detail.Detail, error)
	CancelPipeline(
		project int64,
		identifier int64,
	) (*pipeline_detail.Detail, error)
	RetryPipeline(
		project int64,
		identifier int64,
	) (*pipeline_detail.Detail, error)
	File(
		project int64,
		branch string,
		name string,
	) (*file.File, error)
	MustFile(
		project int64,
		branch string,
		name string,
	) *file.File
	Tree(
		project int64,
		path string,
		reference string,
		recursive bool,
		limit int64,
	) ([]*tree.Node, error)
	CreateBranch(
		project int64,
		name string,
		reference string,
	) (*branch.Branch, error)
	CommitActions(
		project int64,
		branch string,
		message string,
		v []*gitlab.CommitActionOptions,
	) (*commit.Commit, error)
	MustCommitActions(
		project int64,
		branch string,
		message string,
		v []*gitlab.CommitActionOptions,
	) *commit.Commit
	RegistryRepositories(
		project int64,
		panicOnForbidden bool,
	) ([]*registry_repository.Repository, error)
	DeleteRegistryRepository(
		project int64,
		repository int64,
	) error
	Images(
		project int64,
		repository int64,
	) ([]*image.Image, error)
}
