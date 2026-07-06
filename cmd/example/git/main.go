package main

import (
	git "github.com/funtimecoding/soil/pkg/git/example"
	github "github.com/funtimecoding/soil/pkg/github/example"
	"github.com/funtimecoding/soil/pkg/gitlab/check/clean_job"
	gitlab "github.com/funtimecoding/soil/pkg/gitlab/example"
	"github.com/funtimecoding/soil/pkg/gitlab/example/graph_query"
)

func main() {
	github.LatestTags()

	if false {
		gitlab.Registry()
		git.BuildInformation()
		clean_job.Check()
		gitlab.BranchRequest()
		gitlab.CloneAll()
		gitlab.Feature()
		gitlab.File()
		graph_query.Query()
		gitlab.Issue()
		gitlab.MergeRequest()
		gitlab.Project()
		gitlab.Runner()
		gitlab.Search()
		github.BranchRequest()
		github.CleanJob()
		github.Packages()
		github.PullRequest()
		github.Search()
	}
}
