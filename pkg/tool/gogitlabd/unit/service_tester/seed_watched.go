package service_tester

import (
	gitlabConstant "github.com/funtimecoding/soil/pkg/gitlab/constant"
	"github.com/funtimecoding/soil/pkg/tool/gogitlabd/constant"
	"gitlab.com/gitlab-org/api/client-go/v3"
	"time"
)

func (o *Tester) SeedWatched() *gitlab.MergeRequest {
	o.t.Helper()
	at := time.Now()
	result := &gitlab.MergeRequest{
		BasicMergeRequest: gitlab.BasicMergeRequest{
			ProjectID:           constant.FixtureProject,
			IID:                 constant.FixtureRequest,
			Title:               "watched change",
			State:               gitlabConstant.OpenedState,
			DetailedMergeStatus: "mergeable",
			SHA:                 "aaaa1111bbbb2222",
			UpdatedAt:           &at,
			References: &gitlab.IssueReferences{
				Full: constant.FixtureReference,
			},
		},
		HeadPipeline: &gitlab.Pipeline{ID: 10, Status: "running"},
	}
	o.Alfa().SeedDetail(result)

	return result
}
