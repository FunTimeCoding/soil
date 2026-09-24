package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/gitlab/merge_request"
	"github.com/funtimecoding/soil/pkg/tool/gogitlabd/worker"
	"testing"
)

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
