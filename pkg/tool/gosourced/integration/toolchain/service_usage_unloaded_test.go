package toolchain

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/errors/load_failure"
	"github.com/funtimecoding/soil/pkg/system/system_tester"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/integration/toolchain/usage_tester"
	"testing"
)

func TestUsageSkipsBrokenCallerPackage(t *testing.T) {
	o := usage_tester.Prepare(t)
	system_tester.Write(
		t,
		o.Directory,
		"pkg/worker/broken.go",
		"package worker\n\nfunc broken() {\n\tundefinedCall()\n}\n",
	)
	c, e := o.Usage.Usage(o.Directory, "example/pkg/service", "Service")
	assert.FatalOnError(t, e)
	assert.Count(t, 1, c.Unloaded)
	assert.String(t, "example/pkg/worker", c.Unloaded[0].Package)
	assert.String(t, "callers", c.Unloaded[0].Lost)
	assert.StringContains(t, "undefinedCall", c.Unloaded[0].Cause)
	place := usage_tester.Method(t, c, "Place").Callers
	assert.Count(t, 2, place)
	assert.String(
		t,
		"production pkg/model_context Place 1",
		usage_tester.CallerLine(place[0]),
	)
	assert.String(
		t,
		"test pkg/unit TestPlace 1",
		usage_tester.CallerLine(place[1]),
	)
}

func TestUsageSkipsBrokenStorePackage(t *testing.T) {
	o := usage_tester.Prepare(t)
	system_tester.Write(
		t,
		o.Directory,
		"pkg/store/broken.go",
		"package store\n\nfunc broken() {\n\tundefinedCall()\n}\n",
	)
	c, e := o.Usage.Usage(o.Directory, "example/pkg/service", "Service")
	assert.FatalOnError(t, e)
	assert.Count(t, 1, c.Unloaded)
	assert.String(t, "example/pkg/store", c.Unloaded[0].Package)
	assert.String(t, "tables", c.Unloaded[0].Lost)
	assert.Count(t, 3, usage_tester.Method(t, c, "Place").Callers)
}

func TestUsageSkipsBrokenModelPackage(t *testing.T) {
	o := usage_tester.Prepare(t)
	system_tester.Write(
		t,
		o.Directory,
		"pkg/model/receipt/broken.go",
		"package receipt\n\nfunc broken() {\n\tundefinedCall()\n}\n",
	)
	c, e := o.Usage.Usage(o.Directory, "example/pkg/service", "Service")
	assert.FatalOnError(t, e)
	assert.Count(t, 1, c.Unloaded)
	assert.String(t, "example/pkg/model/receipt", c.Unloaded[0].Package)
	assert.String(t, "tables", c.Unloaded[0].Lost)
}

func TestUsageFailsOnBrokenTypePackage(t *testing.T) {
	o := usage_tester.Prepare(t)
	system_tester.Write(
		t,
		o.Directory,
		"pkg/service/broken.go",
		"package service\n\nfunc broken() {\n\tundefinedCall()\n}\n",
	)
	_, e := o.Usage.Usage(o.Directory, "example/pkg/service", "Service")
	failure, okay := load_failure.As(e)
	assert.True(t, okay)
	assert.String(t, "example/pkg/service", failure.Package)
}

func TestUsageSkipsBrokenTestOfTypePackage(t *testing.T) {
	o := usage_tester.Prepare(t)
	system_tester.Write(
		t,
		o.Directory,
		"pkg/service/broken_test.go",
		"package service\n\nimport \"testing\"\n\nfunc TestBroken(t *testing.T) {\n\tundefinedCall()\n}\n",
	)
	c, e := o.Usage.Usage(o.Directory, "example/pkg/service", "Service")
	assert.FatalOnError(t, e)
	assert.Count(t, 1, c.Unloaded)
	assert.String(
		t,
		"example/pkg/service [example/pkg/service.test]",
		c.Unloaded[0].Package,
	)
	assert.String(t, "callers", c.Unloaded[0].Lost)
}
