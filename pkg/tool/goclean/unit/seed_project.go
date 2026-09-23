package unit

import (
	"github.com/funtimecoding/soil/pkg/git/constant"
	constant1 "github.com/funtimecoding/soil/pkg/gitlab/constant"
	"github.com/funtimecoding/soil/pkg/gitlab/mock_client"
	constant2 "github.com/funtimecoding/soil/pkg/tool/goclean/constant"
)

func seedProject() *mock_client.Client {
	c := mock_client.New()
	c.SeedBranch(constant.MainBranch, constant2.FixtureMainHash)
	c.SeedPipeline(
		1,
		constant.MainBranch,
		constant2.FixtureMainHash,
		constant1.JobSuccess,
	)
	c.SeedPipeline(
		2,
		constant2.FixtureReleaseName,
		constant2.FixtureMainHash,
		constant1.JobSuccess,
	)
	c.SeedPipeline(
		3,
		constant2.FixtureFeatureName,
		constant2.FixtureReleaseHash,
		constant1.JobSuccess,
	)
	c.SeedPipeline(
		4,
		constant2.FixtureFeatureName,
		constant2.FixtureBuildHash,
		constant1.JobRunning,
	)
	c.SeedPipeline(
		5,
		constant2.FixtureFeatureName,
		constant2.FixtureQueuedHash,
		constant1.JobPending,
	)

	return c
}
