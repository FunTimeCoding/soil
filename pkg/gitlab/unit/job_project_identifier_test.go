package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/gitlab/job"
	"gitlab.com/gitlab-org/api/client-go/v3"
	"testing"
)

func TestJobProjectIdentifierFromProject(t *testing.T) {
	assert.Integer(
		t,
		int64(5),
		job.New(
			&gitlab.Job{
				Project:  &gitlab.Project{ID: 5},
				Pipeline: gitlab.JobPipeline{ProjectID: 12},
			},
		).ProjectIdentifier(),
	)
}

func TestJobProjectIdentifierFallsBackToPipeline(t *testing.T) {
	assert.Integer(
		t,
		int64(12),
		job.New(
			&gitlab.Job{
				Project:  &gitlab.Project{},
				Pipeline: gitlab.JobPipeline{ProjectID: 12},
			},
		).ProjectIdentifier(),
	)
}

func TestJobProjectIdentifierWithoutProject(t *testing.T) {
	assert.Integer(t, int64(0), job.New(&gitlab.Job{}).ProjectIdentifier())
}
