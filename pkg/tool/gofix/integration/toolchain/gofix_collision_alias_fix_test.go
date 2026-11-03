package toolchain

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/lint/analyzer/testutil"
	"github.com/funtimecoding/soil/pkg/lint/output"
	"github.com/funtimecoding/soil/pkg/tool/gofix"
	"github.com/funtimecoding/soil/pkg/tool/gofix/integration/module_tester"
	"path/filepath"
	"testing"
)

func TestCollisionAliasFixAppliesEveryFirstPartyCollision(t *testing.T) {
	directory := module_tester.CollisionKinds(t)
	r := output.NewResultsWithDirectory(directory)
	gofix.RunCollisionAliasFixWithDirectory(
		[]string{"./..."},
		directory,
		nil,
		false,
		r,
	)
	t.Run(
		"KeeperTakesThePlainName",
		func(t *testing.T) {
			assert.String(
				t,
				"package service\n\nimport (\n\ttestmoduleFace \"testmodule/pkg/face\"\n\t\"testmodule/pkg/golf/face\"\n)\n\nfunc Service(testmoduleFace.Reporter, face.Client) {}\n",
				testutil.ReadFile(
					t,
					filepath.Join(directory, "pkg/golf/service/service.go"),
				),
			)
		},
	)
	t.Run(
		"EveryCollidingMockCarriesItsOwner",
		func(t *testing.T) {
			assert.String(
				t,
				"package unit\n\nimport (\n\thotelMock \"testmodule/pkg/hotel/mock_client\"\n\tgoindiaMock \"testmodule/pkg/tool/goindia/client/mock_client\"\n)\n\nfunc Doubles() (*hotelMock.Client, *goindiaMock.Client) {\n\treturn &hotelMock.Client{}, &goindiaMock.Client{}\n}\n",
				testutil.ReadFile(
					t,
					filepath.Join(directory, "pkg/tool/goindia/unit/doubles.go"),
				),
			)
		},
	)
	t.Run(
		"AHeldGroupStillHolds",
		func(t *testing.T) {
			assert.String(
				t,
				"package held\n\nimport (\n\t\"testmodule\"\n\tchain1 \"testmodule/pkg/delta/chain\"\n)\n\nfunc Held(*chain.Root) *chain1.Delta {\n\treturn &chain1.Delta{}\n}\n",
				testutil.ReadFile(
					t,
					filepath.Join(directory, "pkg/delta/held/held.go"),
				),
			)
		},
	)
	t.Run(
		"ModuleStillBuilds",
		func(t *testing.T) {
			assertBuilds(t, directory)
		},
	)
}
