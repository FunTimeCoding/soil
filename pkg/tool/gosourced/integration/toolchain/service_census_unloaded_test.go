package toolchain

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/errors/load_failure"
	"github.com/funtimecoding/soil/pkg/lint/analyzer/testutil"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/unit/service_tester"
	"testing"
)

func TestCensusSkipsBrokenCallerPackage(t *testing.T) {
	d := prepareCensus(t)
	testutil.WriteFile(
		t,
		d,
		"pkg/worker/broken.go",
		"package worker\n\nfunc broken() {\n\tundefinedCall()\n}\n",
	)
	c, e := service_tester.NewService(t).Census.Census(
		d,
		"example/pkg/service",
		"Service",
	)
	assert.FatalOnError(t, e)
	assert.Count(t, 1, c.Unloaded)
	assert.String(t, "example/pkg/worker", c.Unloaded[0].Package)
	assert.String(t, "callers", c.Unloaded[0].Lost)
	assert.StringContains(t, "undefinedCall", c.Unloaded[0].Cause)
	place := censusMethod(t, c, "Place").Callers
	assert.Count(t, 2, place)
	assert.String(
		t,
		"production pkg/model_context Place 1",
		callerLine(place[0]),
	)
	assert.String(t, "test pkg/unit TestPlace 1", callerLine(place[1]))
}

func TestCensusSkipsBrokenStorePackage(t *testing.T) {
	d := prepareCensus(t)
	testutil.WriteFile(
		t,
		d,
		"pkg/store/broken.go",
		"package store\n\nfunc broken() {\n\tundefinedCall()\n}\n",
	)
	c, e := service_tester.NewService(t).Census.Census(
		d,
		"example/pkg/service",
		"Service",
	)
	assert.FatalOnError(t, e)
	assert.Count(t, 1, c.Unloaded)
	assert.String(t, "example/pkg/store", c.Unloaded[0].Package)
	assert.String(t, "tables", c.Unloaded[0].Lost)
	assert.Count(t, 3, censusMethod(t, c, "Place").Callers)
}

func TestCensusSkipsBrokenModelPackage(t *testing.T) {
	d := prepareCensus(t)
	testutil.WriteFile(
		t,
		d,
		"pkg/model/receipt/broken.go",
		"package receipt\n\nfunc broken() {\n\tundefinedCall()\n}\n",
	)
	c, e := service_tester.NewService(t).Census.Census(
		d,
		"example/pkg/service",
		"Service",
	)
	assert.FatalOnError(t, e)
	assert.Count(t, 1, c.Unloaded)
	assert.String(t, "example/pkg/model/receipt", c.Unloaded[0].Package)
	assert.String(t, "tables", c.Unloaded[0].Lost)
}

func TestCensusRefusesBrokenTypePackage(t *testing.T) {
	d := prepareCensus(t)
	testutil.WriteFile(
		t,
		d,
		"pkg/service/broken.go",
		"package service\n\nfunc broken() {\n\tundefinedCall()\n}\n",
	)
	_, e := service_tester.NewService(t).Census.Census(
		d,
		"example/pkg/service",
		"Service",
	)
	failure, okay := load_failure.As(e)
	assert.True(t, okay)
	assert.String(t, "example/pkg/service", failure.Package)
}

func TestCensusSkipsBrokenTestOfTypePackage(t *testing.T) {
	d := prepareCensus(t)
	testutil.WriteFile(
		t,
		d,
		"pkg/service/broken_test.go",
		"package service\n\nimport \"testing\"\n\nfunc TestBroken(t *testing.T) {\n\tundefinedCall()\n}\n",
	)
	c, e := service_tester.NewService(t).Census.Census(
		d,
		"example/pkg/service",
		"Service",
	)
	assert.FatalOnError(t, e)
	assert.Count(t, 1, c.Unloaded)
	assert.String(
		t,
		"example/pkg/service [example/pkg/service.test]",
		c.Unloaded[0].Package,
	)
	assert.String(t, "callers", c.Unloaded[0].Lost)
}
