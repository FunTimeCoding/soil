package unit_test

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/gitlab/branch"
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

func pipeline(
	identifier int64,
	reference string,
	status string,
) *gitlab.PipelineInfo {
	return &gitlab.PipelineInfo{
		ID:     identifier,
		Ref:    reference,
		Status: status,
	}
}

func TestLatestPipelineNewestWins(t *testing.T) {
	r := latest_pipeline.New(
		branches("main"),
		[]*gitlab.PipelineInfo{
			pipeline(1, "main", "failed"),
			pipeline(3, "main", "success"),
			pipeline(2, "main", "failed"),
		},
	)
	assert.Count(t, 1, r)
	assert.Integer(t, int64(3), r["main"].ID)
	assert.String(t, "success", r["main"].Status)
}

func TestLatestPipelineSkipsDeletedBranch(t *testing.T) {
	r := latest_pipeline.New(
		branches("main"),
		[]*gitlab.PipelineInfo{
			pipeline(1, "main", "success"),
			pipeline(2, "renovate/gone", "failed"),
		},
	)
	assert.Count(t, 1, r)
	assert.NotNil(t, r["main"])
	assert.Nil(t, r["renovate/gone"])
}

func TestLatestPipelineSeparatesReferences(t *testing.T) {
	r := latest_pipeline.New(
		branches("main", "topic"),
		[]*gitlab.PipelineInfo{
			pipeline(1, "main", "success"),
			pipeline(2, "topic", "failed"),
		},
	)
	assert.Count(t, 2, r)
	assert.String(t, "success", r["main"].Status)
	assert.String(t, "failed", r["topic"].Status)
}

func TestLatestPipelineWithoutPipelines(t *testing.T) {
	assert.Count(t, 0, latest_pipeline.New(branches("main"), nil))
}

func TestLatestPipelineWithoutBranches(t *testing.T) {
	assert.Count(
		t,
		0,
		latest_pipeline.New(
			nil,
			[]*gitlab.PipelineInfo{pipeline(1, "main", "success")},
		),
	)
}
