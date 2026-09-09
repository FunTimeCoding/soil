package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/gitlab/branch"
	"github.com/funtimecoding/soil/pkg/gitlab/pipeline"
	"github.com/funtimecoding/soil/pkg/gitlab/tag"
	"github.com/funtimecoding/soil/pkg/tool/gogitlabd/types/latest_pipeline"
	"gitlab.com/gitlab-org/api/client-go/v2"
	"testing"
)

func branches(name ...string) []*branch.Branch {
	var result []*branch.Branch

	for _, n := range name {
		result = append(
			result,
			branch.New(&gitlab.Branch{Name: n, Commit: &gitlab.Commit{}}),
		)
	}

	return result
}

func tags(name ...string) []*tag.Tag {
	var result []*tag.Tag

	for _, n := range name {
		result = append(result, tag.New(&gitlab.Tag{Name: n}))
	}

	return result
}

func newPipeline(
	identifier int64,
	reference string,
	status string,
) *pipeline.Pipeline {
	return pipeline.New(
		&gitlab.PipelineInfo{
			ID:     identifier,
			Ref:    reference,
			Status: status,
		},
	)
}

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
