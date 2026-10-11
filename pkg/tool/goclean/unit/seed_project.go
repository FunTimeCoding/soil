package unit

import (
	gitConstant "github.com/funtimecoding/soil/pkg/git/constant"
	gitlabConstant "github.com/funtimecoding/soil/pkg/gitlab/constant"
	"github.com/funtimecoding/soil/pkg/gitlab/mock_client"
	"github.com/funtimecoding/soil/pkg/tool/goclean/constant"
)

func seedProject() *mock_client.Client {
	c := mock_client.New()
	c.SeedBranch(gitConstant.MainBranch, constant.FixtureMainHash)
	c.SeedPipeline(
		1,
		gitConstant.MainBranch,
		constant.FixtureMainHash,
		gitlabConstant.JobSuccess,
	)
	c.SeedPipeline(
		2,
		constant.FixtureReleaseName,
		constant.FixtureMainHash,
		gitlabConstant.JobSuccess,
	)
	c.SeedPipeline(
		3,
		constant.FixtureFeatureName,
		constant.FixtureReleaseHash,
		gitlabConstant.JobSuccess,
	)
	c.SeedPipeline(
		4,
		constant.FixtureFeatureName,
		constant.FixtureBuildHash,
		gitlabConstant.JobRunning,
	)
	c.SeedPipeline(
		5,
		constant.FixtureFeatureName,
		constant.FixtureQueuedHash,
		gitlabConstant.JobPending,
	)

	return c
}
