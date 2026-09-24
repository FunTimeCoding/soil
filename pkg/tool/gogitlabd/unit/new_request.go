package unit

import (
	"github.com/funtimecoding/soil/pkg/gitlab/merge_request"
	"gitlab.com/gitlab-org/api/client-go/v3"
	"time"
)

func newRequest(
	project int64,
	identifier int64,
	title string,
	createdDaysAgo int,
) *merge_request.Request {
	created := time.Now().AddDate(0, 0, -createdDaysAgo)

	return merge_request.New(
		&gitlab.BasicMergeRequest{
			ProjectID: project,
			IID:       identifier,
			Title:     title,
			State:     "opened",
			CreatedAt: &created,
		},
	)
}
