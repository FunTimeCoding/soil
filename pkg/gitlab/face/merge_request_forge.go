package face

import (
	"github.com/funtimecoding/soil/pkg/gitlab/commit"
	"github.com/funtimecoding/soil/pkg/gitlab/diff"
	"github.com/funtimecoding/soil/pkg/gitlab/discussion"
	"github.com/funtimecoding/soil/pkg/gitlab/merge_request"
	"github.com/funtimecoding/soil/pkg/gitlab/merge_request_detail"
	"github.com/funtimecoding/soil/pkg/gitlab/note"
)

type MergeRequestForge interface {
	ProjectMergeRequests(
		project int64,
		state string,
	) ([]*merge_request.Request, error)
	AssignedMergeRequests(all bool) ([]*merge_request.Request, error)
	ReviewingMergeRequests(all bool) ([]*merge_request.Request, error)
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
	AssignedReviews(all bool) ([]*merge_request.Request, error)
	BranchRequest(
		project int64,
		branch string,
	) (*merge_request.Request, error)
	FindBranchRequest(
		project int64,
		branch string,
	) (*merge_request.Request, bool, error)
	MergeRequestCommits(
		project int64,
		request int64,
	) ([]*commit.Commit, error)
	MustBranchRequest(
		project int64,
		branch string,
	) *merge_request.Request
	MustProjectsMergeRequests() []*merge_request.Request
	ProjectsMergeRequests() ([]*merge_request.Request, error)
}
