package toolchain

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/lint/analyzer/testutil"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/integration/toolchain/rewrite_tester"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/unit/service_tester"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestExtractType(t *testing.T) {
	o := service_tester.New(t, service_tester.WithFixture("extract-basic/src"))
	r, e := o.Symbol.ExtractType(
		t.Context(),
		o.Directory,
		"example/pkg/target",
		"Store",
		"example/pkg/model",
		"",
		true,
		false,
	)
	assert.FatalOnError(t, e)
	testutil.AssertBlocked(t, r, 0)
	_, e = os.Stat(filepath.Join(o.Directory, "pkg/target/store.go"))
	assert.True(t, os.IsNotExist(e))
	_, e = os.Stat(filepath.Join(o.Directory, "pkg/target/save.go"))
	assert.True(t, os.IsNotExist(e))
	store := o.Read("pkg/model/store.go")
	rewrite_tester.Formatted(t, store)
	assert.StringContains(t, "package model", store)
	assert.StringContains(t, "Name string", store)
	save := o.Read("pkg/model/save.go")
	rewrite_tester.Formatted(t, save)
	assert.StringContains(t, "func (s *Store) Save() string", save)
	assert.StringContains(t, "return s.Name", save)
	run := o.Read("pkg/target/run.go")
	rewrite_tester.Formatted(t, run)
	assert.StringContains(t, "v := model.Store{Name: \"alfa\"}", run)
	assert.StringContains(t, "return v.Save()", run)
	caller := o.Read("pkg/caller/run.go")
	rewrite_tester.Formatted(t, caller)
	assert.StringContains(t, "func Run() *model.Store", caller)
	assert.StringContains(t, "return &model.Store{}", caller)
	assert.False(t, strings.Contains(caller, "target"))
}

func TestExtractTypeIntoNamesakePackageImportsPlain(t *testing.T) {
	o := service_tester.New(
		t,
		service_tester.WithFixture("extract-namesake/src"),
	)
	r, e := o.Symbol.ExtractType(
		t.Context(),
		o.Directory,
		"example/pkg/alfa",
		"waiter",
		"example/pkg/alfa/waiter",
		"",
		true,
		false,
	)
	assert.FatalOnError(t, e)
	testutil.AssertBlocked(t, r, 0)
	queue := o.Read("pkg/alfa/queue.go")
	assert.StringContains(t, "import \"example/pkg/alfa/waiter\"\n", queue)
	assert.StringContains(t, "[]*waiter.Waiter", queue)
	rewrite_tester.Compiles(t, o.Directory)
}

func TestExtractTypeInternalStaysUnexported(t *testing.T) {
	o := service_tester.New(
		t,
		service_tester.WithFixture("extract-internal/src"),
	)
	r, e := o.Symbol.ExtractType(
		t.Context(),
		o.Directory,
		"example/pkg/target",
		"Store",
		"example/pkg/model",
		"",
		true,
		false,
	)
	assert.FatalOnError(t, e)
	testutil.AssertBlocked(t, r, 0)
	helper := o.Read("pkg/model/helper.go")
	rewrite_tester.Formatted(t, helper)
	assert.StringContains(t, "func (s *Store) helper() string", helper)
	describe := o.Read("pkg/model/describe.go")
	assert.StringContains(t, "return s.helper()", describe)
	run := o.Read("pkg/target/run.go")
	assert.StringContains(t, "v := model.Store{}", run)
}

func TestExtractTypeCollision(t *testing.T) {
	o := service_tester.New(
		t,
		service_tester.WithFixture("extract-collision/src"),
	)
	r, e := o.Symbol.ExtractType(
		t.Context(),
		o.Directory,
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

func TestExtractTypeQualifiesBackReference(t *testing.T) {
	o := service_tester.New(
		t,
		service_tester.WithFixture("extract-dependency/src"),
	)
	r, e := o.Symbol.ExtractType(
		t.Context(),
		o.Directory,
		"example/pkg/target",
		"Store",
		"example/pkg/model",
		"",
		true,
		false,
	)
	assert.FatalOnError(t, e)
	testutil.AssertBlocked(t, r, 0)
	rewrite_tester.FixedContains(t, "prefix → Prefix", r)
	rewrite_tester.FixedContains(
		t,
		"prefix → target.Prefix (back reference)",
		r,
	)
	assert.StringContains(
		t,
		"return target.Prefix()",
		o.Read("pkg/model/describe.go"),
	)
	assert.StringContains(
		t,
		"func Prefix() string",
		o.Read("pkg/target/prefix.go"),
	)
	rewrite_tester.Compiles(t, o.Directory)
}

func TestExtractTypeBackReferenceCycleRefused(t *testing.T) {
	o := service_tester.New(
		t,
		service_tester.WithFixture("extract-dependency-cycle/src"),
	)
	r, e := o.Symbol.ExtractType(
		t.Context(),
		o.Directory,
		"example/pkg/target",
		"Store",
		"example/pkg/model",
		"",
		true,
		false,
	)
	assert.FatalOnError(t, e)
	testutil.AssertBlocked(t, r, 1)
	testutil.AssertBlockedContains(
		t,
		r,
		"example/pkg/model → example/pkg/target → example/pkg/model",
	)
}

func TestExtractTypeFieldCollision(t *testing.T) {
	o := service_tester.New(
		t,
		service_tester.WithFixture("extract-field-collision/src"),
	)
	r, e := o.Symbol.ExtractType(
		t.Context(),
		o.Directory,
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
	o := service_tester.New(
		t,
		service_tester.WithFixture("extract-unexported-type/src"),
	)
	r, e := o.Symbol.ExtractType(
		t.Context(),
		o.Directory,
		"example/pkg/target",
		"store",
		"example/pkg/model",
		"",
		true,
		false,
	)
	assert.FatalOnError(t, e)
	testutil.AssertBlocked(t, r, 0)
	moved := o.Read("pkg/model/store.go")
	rewrite_tester.Formatted(t, moved)
	assert.StringContains(t, "type Store struct{}", moved)
	factory := o.Read("pkg/target/new.go")
	rewrite_tester.Formatted(t, factory)
	assert.StringContains(t, "func New() *model.Store", factory)
	assert.StringContains(t, "return &model.Store{}", factory)
}

func TestExtractTypeTargetFile(t *testing.T) {
	o := service_tester.New(
		t,
		service_tester.WithFixture("extract-internal/src"),
	)
	r, e := o.Symbol.ExtractType(
		t.Context(),
		o.Directory,
		"example/pkg/target",
		"Store",
		"example/pkg/model",
		"store.go",
		true,
		false,
	)
	assert.FatalOnError(t, e)
	testutil.AssertBlocked(t, r, 0)
	moved := o.Read("pkg/model/store.go")
	rewrite_tester.Formatted(t, moved)
	assert.StringContains(t, "type Store struct{}", moved)
	assert.StringContains(t, "func (s *Store) helper() string", moved)
	assert.StringContains(t, "func (s *Store) Describe() string", moved)
	_, e = os.Stat(filepath.Join(o.Directory, "pkg/model/helper.go"))
	assert.True(t, os.IsNotExist(e))
}
