package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/lint/analyzer/testutil"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/unit/service_tester"
	"testing"
)

func TestRenameField(t *testing.T) {
	d := testutil.PrepareTestPackage(
		t,
		service_tester.ServiceTestdata("rename-field/src"),
	)
	r, e := testService().Rename(
		d,
		"example/pkg/target",
		"api",
		"requester",
		"Store",
		false,
	)
	assert.FatalOnError(t, e)
	testutil.AssertBlocked(t, r, 0)
	store := service_tester.ReadFixtureFile(t, d, "pkg/target/store.go")
	assert.StringContains(t, "requester string", store)
	constructor := service_tester.ReadFixtureFile(t, d, "pkg/target/new.go")
	assert.StringContains(t, "func New(api string)", constructor)
	assert.StringContains(t, "&Store{requester: api}", constructor)
	describe := service_tester.ReadFixtureFile(t, d, "pkg/target/describe.go")
	assert.StringContains(t, "s.requester + s.Name", describe)
}

func TestRenameExportedFieldAcrossPackages(t *testing.T) {
	d := testutil.PrepareTestPackage(
		t,
		service_tester.ServiceTestdata("rename-field/src"),
	)
	r, e := testService().Rename(
		d,
		"example/pkg/target",
		"Name",
		"Title",
		"Store",
		false,
	)
	assert.FatalOnError(t, e)
	testutil.AssertBlocked(t, r, 0)
	describe := service_tester.ReadFixtureFile(t, d, "pkg/target/describe.go")
	assert.StringContains(t, "s.api + s.Title", describe)
	caller := service_tester.ReadFixtureFile(t, d, "pkg/caller/run.go")
	assert.StringContains(t, "&target.Store{Title: \"alfa\"}", caller)
	assert.StringContains(t, "return s.Title", caller)
}

func TestRenameFieldCollidesWithField(t *testing.T) {
	d := testutil.PrepareTestPackage(
		t,
		service_tester.ServiceTestdata("rename-field/src"),
	)
	r, e := testService().Rename(
		d,
		"example/pkg/target",
		"api",
		"Name",
		"Store",
		false,
	)
	assert.FatalOnError(t, e)
	testutil.AssertBlocked(t, r, 1)
	testutil.AssertBlockedContains(t, r, "field Name already exists")
}

func TestRenameFieldCollidesWithMethod(t *testing.T) {
	d := testutil.PrepareTestPackage(
		t,
		service_tester.ServiceTestdata("rename-field/src"),
	)
	r, e := testService().Rename(
		d,
		"example/pkg/target",
		"api",
		"Describe",
		"Store",
		false,
	)
	assert.FatalOnError(t, e)
	testutil.AssertBlocked(t, r, 1)
	testutil.AssertBlockedContains(t, r, "method Describe already exists")
}

func TestRenameEmbeddedFieldRefused(t *testing.T) {
	d := testutil.PrepareTestPackage(
		t,
		service_tester.ServiceTestdata("rename-field/src"),
	)
	r, e := testService().Rename(
		d,
		"example/pkg/target",
		"Base",
		"Root",
		"Store",
		false,
	)
	assert.FatalOnError(t, e)
	testutil.AssertBlocked(t, r, 1)
	testutil.AssertBlockedContains(t, r, "embedded field")
	store := service_tester.ReadFixtureFile(t, d, "pkg/target/store.go")
	assert.StringContains(t, "\tBase\n", store)
}

func TestRenameFieldToUnexportedBlockedByCrossPackage(t *testing.T) {
	d := testutil.PrepareTestPackage(
		t,
		service_tester.ServiceTestdata("rename-field/src"),
	)
	r, e := testService().Rename(
		d,
		"example/pkg/target",
		"Name",
		"name",
		"Store",
		false,
	)
	assert.FatalOnError(t, e)
	testutil.AssertBlockedContains(t, r, "would lose access")
}

func TestFindReferencesField(t *testing.T) {
	d := testutil.PrepareTestPackage(
		t,
		service_tester.ServiceTestdata("rename-field/src"),
	)
	r, references, e := testService().FindReferences(
		d,
		"example/pkg/target",
		"api",
		"Store",
	)
	assert.FatalOnError(t, e)
	testutil.AssertBlocked(t, r, 0)
	assert.Integer(t, 2, references.Total)
	assert.String(t, "pkg/target/describe.go", references.Locations[0].File)
	assert.String(t, "pkg/target/new.go", references.Locations[1].File)
}

func TestChangeVisibilityField(t *testing.T) {
	d := testutil.PrepareTestPackage(
		t,
		service_tester.ServiceTestdata("rename-field/src"),
	)
	r, e := testService().ChangeVisibility(
		d,
		"api",
		"example/pkg/target",
		"Store",
		false,
	)
	assert.FatalOnError(t, e)
	testutil.AssertBlocked(t, r, 0)
	constructor := service_tester.ReadFixtureFile(t, d, "pkg/target/new.go")
	assert.StringContains(t, "&Store{Api: api}", constructor)
}
