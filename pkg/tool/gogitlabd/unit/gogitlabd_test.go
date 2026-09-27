package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/gitlab/job"
	"github.com/funtimecoding/soil/pkg/gitlab/merge_request"
	"github.com/funtimecoding/soil/pkg/gitlab/pipeline"
	"github.com/funtimecoding/soil/pkg/notation"
	"github.com/funtimecoding/soil/pkg/tool/gogitlabd/model_context/response"
	"github.com/funtimecoding/soil/pkg/tool/gogitlabd/types/latest_pipeline"
	"github.com/funtimecoding/soil/pkg/tool/gogitlabd/worker"
	"gitlab.com/gitlab-org/api/client-go/v3"
	"testing"
)

func TestLatestPipelineNewestWins(t *testing.T) {
	r := latest_pipeline.New(
		branches("main"),
		nil,
		[]*pipeline.Pipeline{
			newPipeline(1, "main", "failed"),
			newPipeline(3, "main", "success"),
			newPipeline(2, "main", "failed"),
		},
	)
	assert.Count(t, 1, r)
	assert.Integer(t, int64(3), r["main"].Identifier)
	assert.String(t, "success", r["main"].Status)
}

func TestLatestPipelineSkipsDeletedBranch(t *testing.T) {
	r := latest_pipeline.New(
		branches("main"),
		nil,
		[]*pipeline.Pipeline{
			newPipeline(1, "main", "success"),
			newPipeline(2, "renovate/gone", "failed"),
		},
	)
	assert.Count(t, 1, r)
	assert.NotNil(t, r["main"])
	assert.Nil(t, r["renovate/gone"])
}

func TestLatestPipelineSeparatesReferences(t *testing.T) {
	r := latest_pipeline.New(
		branches("main", "topic"),
		nil,
		[]*pipeline.Pipeline{
			newPipeline(1, "main", "success"),
			newPipeline(2, "topic", "failed"),
		},
	)
	assert.Count(t, 2, r)
	assert.String(t, "success", r["main"].Status)
	assert.String(t, "failed", r["topic"].Status)
}

func TestLatestPipelineWithoutPipelines(t *testing.T) {
	assert.Count(t, 0, latest_pipeline.New(branches("main"), nil, nil))
}

func TestLatestPipelineWithoutBranches(t *testing.T) {
	assert.Count(
		t,
		0,
		latest_pipeline.New(
			nil,
			nil,
			[]*pipeline.Pipeline{newPipeline(1, "main", "success")},
		),
	)
}

func TestLatestPipelineIncludesLatestTag(t *testing.T) {
	r := latest_pipeline.New(
		branches("main"),
		tags("v0.2.71", "v0.2.72"),
		[]*pipeline.Pipeline{
			newPipeline(1, "main", "success"),
			newPipeline(2, "v0.2.71", "success"),
			newPipeline(4, "v0.2.72", "running"),
			newPipeline(3, "v0.2.72", "canceled"),
		},
	)
	assert.Count(t, 2, r)
	assert.Integer(t, int64(4), r["v0.2.72"].Identifier)
	assert.String(t, "running", r["v0.2.72"].Status)
	assert.Nil(t, r["v0.2.71"])
}

func TestLatestPipelineSkipsDeletedTag(t *testing.T) {
	r := latest_pipeline.New(
		branches("main"),
		tags("v0.2.72"),
		[]*pipeline.Pipeline{
			newPipeline(1, "main", "success"),
			newPipeline(2, "v0.1.9", "failed"),
		},
	)
	assert.Count(t, 1, r)
	assert.Nil(t, r["v0.1.9"])
}

func TestModelContextResponseJobOmitsProjectAndRaw(t *testing.T) {
	assert.String(
		t,
		"{\n\t\"identifier\": 7,\n\t\"name\": \"generate\",\n\t\"status\": \"success\",\n\t\"stage\": \"publish\",\n\t\"link\": \"https://gitlab.example.org/a/b/-/jobs/7\"\n}",
		notation.MarshalIndent(
			response.NewJob(
				job.New(
					&gitlab.Job{
						ID:     7,
						Name:   "generate",
						Status: "success",
						Stage:  "publish",
						WebURL: "https://gitlab.example.org/a/b/-/jobs/7",
					},
				),
			),
		),
	)
}

func TestMergeRequestsUnionKeepsBothRoles(t *testing.T) {
	r := worker.MergeRequests(
		[]*merge_request.Request{newRequest(1, 10, "assigned", 1)},
		[]*merge_request.Request{newRequest(2, 20, "reviewing", 2)},
	)
	assert.Count(t, 2, r)
}

func TestMergeRequestsDeduplicatesAcrossRoles(t *testing.T) {
	r := worker.MergeRequests(
		[]*merge_request.Request{newRequest(1, 10, "both", 1)},
		[]*merge_request.Request{newRequest(1, 10, "both", 1)},
	)
	assert.Count(t, 1, r)
}

func TestMergeRequestsSeparatesSameIdentifierInDifferentProjects(t *testing.T) {
	r := worker.MergeRequests(
		[]*merge_request.Request{newRequest(1, 10, "first", 1)},
		[]*merge_request.Request{newRequest(2, 10, "second", 2)},
	)
	assert.Count(t, 2, r)
}

func TestMergeRequestsNewestFirst(t *testing.T) {
	r := worker.MergeRequests(
		[]*merge_request.Request{
			newRequest(1, 10, "older", 9),
			newRequest(1, 11, "newest", 1),
		},
		nil,
	)
	assert.String(t, "newest", r[0].Title)
	assert.String(t, "older", r[1].Title)
}

func TestMergeRequestsCapsAtThree(t *testing.T) {
	r := worker.MergeRequests(
		[]*merge_request.Request{
			newRequest(1, 10, "one", 1),
			newRequest(1, 11, "two", 2),
			newRequest(1, 12, "three", 3),
		},
		[]*merge_request.Request{
			newRequest(1, 13, "four", 4),
			newRequest(1, 14, "five", 5),
		},
	)
	assert.Count(t, 3, r)
}
