package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/lint/analyzer/testutil"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/unit/service_tester"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestExtractType(t *testing.T) {
	d := testutil.PrepareTestPackage(
		t,
		service_tester.ServiceTestdata("extract-basic/src"),
	)
	s := testService()
	r, e := s.ExtractType(
		d,
		"example/pkg/target",
		"Store",
		"example/pkg/model",
		"",
		true,
		false,
	)
	assert.FatalOnError(t, e)
	testutil.AssertBlocked(t, r, 0)
	_, e = os.Stat(filepath.Join(d, "pkg/target/store.go"))
	assert.True(t, os.IsNotExist(e))
	_, e = os.Stat(filepath.Join(d, "pkg/target/save.go"))
	assert.True(t, os.IsNotExist(e))
	store := service_tester.ReadFixtureFile(t, d, "pkg/model/store.go")
	assertFormatted(t, store)
	assert.StringContains(t, "package model", store)
	assert.StringContains(t, "Name string", store)
	save := service_tester.ReadFixtureFile(t, d, "pkg/model/save.go")
	assertFormatted(t, save)
	assert.StringContains(t, "func (s *Store) Save() string", save)
	assert.StringContains(t, "return s.Name", save)
	run := service_tester.ReadFixtureFile(t, d, "pkg/target/run.go")
	assertFormatted(t, run)
	assert.StringContains(t, "v := model.Store{Name: \"alfa\"}", run)
	assert.StringContains(t, "return v.Save()", run)
	caller := service_tester.ReadFixtureFile(t, d, "pkg/caller/run.go")
	assertFormatted(t, caller)
	assert.StringContains(t, "func Run() *model.Store", caller)
	assert.StringContains(t, "return &model.Store{}", caller)
	assert.False(t, strings.Contains(caller, "target"))
}

func TestExtractTypeInternalStaysUnexported(t *testing.T) {
	d := testutil.PrepareTestPackage(
		t,
		service_tester.ServiceTestdata("extract-internal/src"),
	)
	s := testService()
	r, e := s.ExtractType(
		d,
		"example/pkg/target",
		"Store",
		"example/pkg/model",
		"",
		true,
		false,
	)
	assert.FatalOnError(t, e)
	testutil.AssertBlocked(t, r, 0)
	helper := service_tester.ReadFixtureFile(t, d, "pkg/model/helper.go")
	assertFormatted(t, helper)
	assert.StringContains(t, "func (s *Store) helper() string", helper)
	describe := service_tester.ReadFixtureFile(t, d, "pkg/model/describe.go")
	assert.StringContains(t, "return s.helper()", describe)
	run := service_tester.ReadFixtureFile(t, d, "pkg/target/run.go")
	assert.StringContains(t, "v := model.Store{}", run)
}

func TestExtractTypeCollision(t *testing.T) {
	d := testutil.PrepareTestPackage(
		t,
		service_tester.ServiceTestdata("extract-collision/src"),
	)
	s := testService()
	r, e := s.ExtractType(
		d,
		"example/pkg/target",
		"Store",
		"example/pkg/model",
		"",
		false,
		false,
	)
	assert.FatalOnError(t, e)
	testutil.AssertBlocked(t, r, 1)
	testutil.AssertBlockedContains(t, r, "already exists")
}

func TestExtractTypeDependency(t *testing.T) {
	d := testutil.PrepareTestPackage(
		t,
		service_tester.ServiceTestdata("extract-dependency/src"),
	)
	s := testService()
	r, e := s.ExtractType(
		d,
		"example/pkg/target",
		"Store",
		"example/pkg/model",
		"",
		true,
		false,
	)
	assert.FatalOnError(t, e)
	testutil.AssertBlocked(t, r, 1)
	testutil.AssertBlockedContains(t, r, "package-local")
	testutil.AssertBlockedContains(t, r, "prefix")
}

func TestExtractTypeFieldCollision(t *testing.T) {
	d := testutil.PrepareTestPackage(
		t,
		service_tester.ServiceTestdata("extract-field-collision/src"),
	)
	s := testService()
	r, e := s.ExtractType(
		d,
		"example/pkg/target",
		"Store",
		"example/pkg/model",
		"",
		true,
		false,
	)
	assert.FatalOnError(t, e)
	testutil.AssertBlocked(t, r, 1)
	testutil.AssertBlockedContains(t, r, "collides")
}

func TestExtractTypeUnexported(t *testing.T) {
	d := testutil.PrepareTestPackage(
		t,
		service_tester.ServiceTestdata("extract-unexported-type/src"),
	)
	s := testService()
	r, e := s.ExtractType(
		d,
		"example/pkg/target",
		"store",
		"example/pkg/model",
		"",
		true,
		false,
	)
	assert.FatalOnError(t, e)
	testutil.AssertBlocked(t, r, 0)
	moved := service_tester.ReadFixtureFile(t, d, "pkg/model/store.go")
	assertFormatted(t, moved)
	assert.StringContains(t, "type Store struct{}", moved)
	factory := service_tester.ReadFixtureFile(t, d, "pkg/target/new.go")
	assertFormatted(t, factory)
	assert.StringContains(t, "func New() *model.Store", factory)
	assert.StringContains(t, "return &model.Store{}", factory)
}

func TestExtractTypeTargetFile(t *testing.T) {
	d := testutil.PrepareTestPackage(
		t,
		service_tester.ServiceTestdata("extract-internal/src"),
	)
	s := testService()
	r, e := s.ExtractType(
		d,
		"example/pkg/target",
		"Store",
		"example/pkg/model",
		"store.go",
		true,
		false,
	)
	assert.FatalOnError(t, e)
	testutil.AssertBlocked(t, r, 0)
	moved := service_tester.ReadFixtureFile(t, d, "pkg/model/store.go")
	assertFormatted(t, moved)
	assert.StringContains(t, "type Store struct{}", moved)
	assert.StringContains(t, "func (s *Store) helper() string", moved)
	assert.StringContains(t, "func (s *Store) Describe() string", moved)
	_, e = os.Stat(filepath.Join(d, "pkg/model/helper.go"))
	assert.True(t, os.IsNotExist(e))
}
