package face

import (
	"context"
	"github.com/funtimecoding/soil/pkg/gitlab/commit"
	"github.com/funtimecoding/soil/pkg/gitlab/diff"
	"github.com/funtimecoding/soil/pkg/gitlab/discussion"
	"github.com/funtimecoding/soil/pkg/gitlab/merge_request"
	"github.com/funtimecoding/soil/pkg/gitlab/merge_request_detail"
	"github.com/funtimecoding/soil/pkg/gitlab/note"
)

type MergeRequestForge interface {
	ProjectMergeRequests(
		x context.Context,
		project int64,
		state string,
	) ([]*merge_request.Request, error)
	AssignedMergeRequests(
		x context.Context,
		all bool,
	) ([]*merge_request.Request, error)
	ReviewingMergeRequests(
		x context.Context,
		all bool,
	) ([]*merge_request.Request, error)
	MergeRequest(
		x context.Context,
		project int64,
		identifier int64,
	) (*merge_request_detail.Detail, error)
	MergeRequestDiffs(
		x context.Context,
		project int64,
		identifier int64,
	) ([]*diff.Diff, error)
	MergeRequestDiscussions(
		x context.Context,
		project int64,
		identifier int64,
	) ([]*discussion.Discussion, error)
	CreateMergeRequestNote(
		x context.Context,
		project int64,
		identifier int64,
		body string,
	) (*note.Note, error)
	AssignedReviews(
		x context.Context,
		all bool,
	) ([]*merge_request.Request, error)
	BranchRequest(
		x context.Context,
		project int64,
		branch string,
	) (*merge_request.Request, error)
	FindBranchRequest(
		x context.Context,
		project int64,
		branch string,
	) (*merge_request.Request, bool, error)
	MergeRequestCommits(
		x context.Context,
		project int64,
		request int64,
	) ([]*commit.Commit, error)
	MustBranchRequest(
		x context.Context,
		project int64,
		branch string,
	) *merge_request.Request
	MustProjectsMergeRequests(x context.Context) []*merge_request.Request
	ProjectsMergeRequests(
		x context.Context,
	) ([]*merge_request.Request, error)
}
