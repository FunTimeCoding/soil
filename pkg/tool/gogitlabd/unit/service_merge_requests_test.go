package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/errors/validation"
	gitlabConstant "github.com/funtimecoding/soil/pkg/gitlab/constant"
	"github.com/funtimecoding/soil/pkg/gitlab/merge_request_filter"
	"github.com/funtimecoding/soil/pkg/tool/gogitlabd/constant"
	"github.com/funtimecoding/soil/pkg/tool/gogitlabd/unit/service_tester"
	"testing"
	"time"
)

func TestMergeRequestsOfLimitOutOfRange(t *testing.T) {
	s := service_tester.New(t).Service()

	for _, limit := range []int{0, 101} {
		f := merge_request_filter.New(7)
		f.Limit = limit
		_, e := s.MergeRequestsOf(t.Context(), constant.FixtureInstance, f)
		assert.True(t, validation.Is(e))
	}
}

func TestMergeRequestsOfNegativeOffsetInvalid(t *testing.T) {
	f := merge_request_filter.New(7)
	f.Offset = -1
	_, e := service_tester.New(t).Service().MergeRequestsOf(
		t.Context(),
		constant.FixtureInstance,
		f,
	)
	assert.True(t, validation.Is(e))
}

func TestMergeRequestsOfUnknownStateInvalid(t *testing.T) {
	f := merge_request_filter.New(7)
	f.State = "open"
	_, e := service_tester.New(t).Service().MergeRequestsOf(
		t.Context(),
		constant.FixtureInstance,
		f,
	)
	assert.True(t, validation.Is(e))
}

func TestMergeRequestsOfUnknownOrderInvalid(t *testing.T) {
	f := merge_request_filter.New(7)
	f.OrderBy = "title"
	_, e := service_tester.New(t).Service().MergeRequestsOf(
		t.Context(),
		constant.FixtureInstance,
		f,
	)
	assert.True(t, validation.Is(e))
}

func TestMergeRequestsOfAllStatesWithFilter(t *testing.T) {
	o := service_tester.New(t)
	now := time.Now()
	o.Alfa().SeedRequest(7, 1, gitlabConstant.OpenedState, "one", "delta", now)
	o.Alfa().SeedRequest(7, 2, gitlabConstant.MergedState, "two", "echo", now)
	o.Alfa().SeedRequest(7, 3, gitlabConstant.ClosedState, "three", "echo", now)
	f := merge_request_filter.New(7)
	f.State = gitlabConstant.AllState
	f.Author = "echo"
	result, e := o.Service().MergeRequestsOf(
		t.Context(),
		constant.FixtureInstance,
		f,
	)
	assert.FatalOnError(t, e)
	assert.Count(t, 2, result)
}
