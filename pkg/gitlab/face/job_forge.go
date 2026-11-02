package face

import (
	"bytes"
	"context"
	"github.com/funtimecoding/soil/pkg/gitlab/job"
	"github.com/funtimecoding/soil/pkg/gitlab/project"
)

type JobForge interface {
	PipelineJob(
		x context.Context,
		project int64,
		identifier int64,
	) (*job.Job, error)
	PipelineJobs(
		x context.Context,
		project int64,
		identifier int64,
	) ([]*job.Job, error)
	CancelJob(
		x context.Context,
		project int64,
		identifier int64,
	) (*job.Job, error)
	Retry(
		x context.Context,
		project int64,
		jobIdentifier int64,
	) (*job.Job, error)
	Trace(
		x context.Context,
		project int64,
		jobIdentifier int64,
	) (string, error)
	ArtifactFile(
		x context.Context,
		project int64,
		job int64,
		path string,
	) (*bytes.Reader, error)
	Artifacts(
		x context.Context,
		project int64,
		job int64,
	) (*bytes.Reader, error)
	ArtifactsFile(
		x context.Context,
		project int64,
		reference string,
		job string,
	) (string, error)
	FailedJobs(x context.Context) ([]*job.Job, error)
	Jobs(x context.Context) ([]*job.Job, error)
	MustJobs(x context.Context) []*job.Job
	MustRetryJob(
		x context.Context,
		j *job.Job,
	) *job.Job
	ProjectJobs(
		x context.Context,
		p *project.Project,
	) ([]*job.Job, error)
	RetryJob(
		x context.Context,
		j *job.Job,
	) (*job.Job, error)
}
