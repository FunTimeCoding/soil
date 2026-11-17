package service_tester

import (
	"github.com/funtimecoding/soil/pkg/errors/sentry/reporter/memory"
	"github.com/funtimecoding/soil/pkg/lint/analyzer/testutil"
	"github.com/funtimecoding/soil/pkg/source/inventory"
	"github.com/funtimecoding/soil/pkg/tool/gomemoryd/client/mock_client"
	"github.com/funtimecoding/soil/pkg/tool/gomemoryd/face"
	"testing"
)

func New(
	t *testing.T,
	o ...Option,
) *Tester {
	t.Helper()
	p := &setup{}

	for _, f := range o {
		f(p)
	}

	result := &Tester{
		Directory: p.directory,
		Memory:    mock_client.New(t),
		Reporter:  memory.New(),
		t:         t,
	}

	if p.fixture != "" {
		result.Directory = testutil.PrepareTestPackage(
			t,
			ServiceTestdata(p.fixture),
			p.modLines...,
		)
	}

	var peer face.Client = result.Memory

	if p.memoryPeer != nil {
		peer = p.memoryPeer
	}

	if p.inventory != nil {
		result.Services = fromParts(t, p.inventory, peer, result.Reporter)

		return result
	}

	result.Services = fromParts(t, inventory.New(), peer, result.Reporter)
	result.Workspace.UseIndex(IndexDirectory())

	return result
}
