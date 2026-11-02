package face

import (
	"bytes"
	"github.com/funtimecoding/soil/pkg/gitlab/job"
	"github.com/funtimecoding/soil/pkg/gitlab/project"
)

type JobForge interface {
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
	ArtifactFile(
		project int64,
		job int64,
		path string,
	) (*bytes.Reader, error)
	Artifacts(
		project int64,
		job int64,
	) (*bytes.Reader, error)
	ArtifactsFile(
		project int64,
		reference string,
		job string,
	) (string, error)
	FailedJobs() ([]*job.Job, error)
	Jobs() ([]*job.Job, error)
	MustJobs() []*job.Job
	MustRetryJob(j *job.Job) *job.Job
	ProjectJobs(p *project.Project) ([]*job.Job, error)
	RetryJob(j *job.Job) (*job.Job, error)
}
