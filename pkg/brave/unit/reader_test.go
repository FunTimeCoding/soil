package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/brave"
	"github.com/funtimecoding/soil/pkg/brave/constant"
	soil "github.com/funtimecoding/soil/pkg/constant"
	"github.com/funtimecoding/soil/pkg/errors/not_found"
	"testing"
)

func TestABookmarkFileIsReadFromItsProfile(t *testing.T) {
	braveSettings(t, constant.FixtureSignedIn, constant.FixtureBookmarkBar)
	b, e := brave.Bookmark(constant.DefaultProfile)
	assert.FatalOnError(t, e)
	assert.String(t, "https://example.org/kid", b.Root.Bar.Children[0].Locator)
}

func TestMissingSettingsAreAnError(t *testing.T) {
	t.Setenv(soil.HomeEnvironment, t.TempDir())
	_, e := brave.Bookmark(constant.DefaultProfile)
	assert.NotNil(t, e)
}

func TestAnUnknownProfileIsNotFound(t *testing.T) {
	braveSettings(t, constant.FixtureSignedIn, constant.FixtureBookmarkBar)
	_, e := brave.ProfileByName("Profile 9")
	assert.True(t, not_found.Is(e))
}

func TestAMissingBookmarkFileIsAnError(t *testing.T) {
	braveSettings(t, constant.FixtureSignedIn, "")
	_, e := brave.Bookmark(constant.DefaultProfile)
	assert.NotNil(t, e)
}

func TestABrokenBookmarkFileIsAnError(t *testing.T) {
	braveSettings(t, constant.FixtureSignedIn, "{roots")
	_, e := brave.Bookmark(constant.DefaultProfile)
	assert.NotNil(t, e)
}

func TestUndecodableAccountsAreAnError(t *testing.T) {
	braveSettings(
		t,
		constant.FixtureBrokenAccounts,
		constant.FixtureBookmarkBar,
	)
	_, e := brave.Profiles()
	assert.NotNil(t, e)
}
