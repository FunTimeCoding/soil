package constant

import "github.com/funtimecoding/soil/pkg/identity"

var Identity = identity.New(
	"goclean",
	"Build artifact cleaner",
	"goclean [flags]",
)

const (
	FixtureMainHash    = "1111111111111111111111111111111111111111"
	FixtureReleaseHash = "2222222222222222222222222222222222222222"
	FixtureBuildHash   = "3333333333333333333333333333333333333333"
	FixtureQueuedHash  = "4444444444444444444444444444444444444444"
	FixtureFeatureName = "feature"
	FixtureReleaseName = "v1.0.0"
)
