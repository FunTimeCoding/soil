package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/brave"
	"github.com/funtimecoding/soil/pkg/brave/constant"
	"github.com/funtimecoding/soil/pkg/brave/unit/brave_tester"
	soil "github.com/funtimecoding/soil/pkg/constant"
	"github.com/funtimecoding/soil/pkg/errors/not_found"
	"testing"
)

func TestBookmarkFileIsReadFromItsProfile(t *testing.T) {
	brave_tester.Settings(
		t,
		constant.FixtureSignedIn,
		constant.FixtureBookmarkBar,
	)
	b, e := brave.Bookmark(constant.DefaultProfile)
	assert.FatalOnError(t, e)
	assert.String(t, "https://example.org/kid", b.Root.Bar.Children[0].Locator)
}

func TestMissingSettingsAreError(t *testing.T) {
	t.Setenv(soil.HomeEnvironment, t.TempDir())
	_, e := brave.Bookmark(constant.DefaultProfile)
	assert.NotNil(t, e)
}

func TestUnknownProfileIsNotFound(t *testing.T) {
	brave_tester.Settings(
		t,
		constant.FixtureSignedIn,
		constant.FixtureBookmarkBar,
	)
	_, e := brave.Profile("Profile 9")
	assert.True(t, not_found.Is(e))
}

func TestMissingBookmarkFileIsError(t *testing.T) {
	brave_tester.Settings(t, constant.FixtureSignedIn, "")
	_, e := brave.Bookmark(constant.DefaultProfile)
	assert.NotNil(t, e)
}

func TestBrokenBookmarkFileIsError(t *testing.T) {
	brave_tester.Settings(t, constant.FixtureSignedIn, "{roots")
	_, e := brave.Bookmark(constant.DefaultProfile)
	assert.NotNil(t, e)
}

func TestUndecodableAccountsAreError(t *testing.T) {
	brave_tester.Settings(
		t,
		constant.FixtureBrokenAccounts,
		constant.FixtureBookmarkBar,
	)
	_, e := brave.Profiles()
	assert.NotNil(t, e)
}
