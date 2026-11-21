package security_tester

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/tool/gosecurityd/store/installed"
	"github.com/funtimecoding/soil/pkg/tool/gosecurityd/store/system"
	vulnerabilityConstant "github.com/funtimecoding/soil/pkg/vulnerability/constant"
	"time"
)

func (o *Tester) ScoredFleet() *Tester {
	o.t.Helper()
	o.UpgradeFleet()
	echo := system.New("echo", "trixie", 1)
	echo.Kernel = "9.12.1-amd64"
	assert.FatalOnError(
		o.t,
		o.System().Put(
			echo,
			[]installed.Installed{
				*installed.New(
					"libtango5",
					"8.4.0-1",
					"tango-utils",
					"8.4.0-1",
					"",
				),
				*installed.New(
					"linux-image-9.12.1-amd64",
					"9.12.1-1",
					"linux-signed-amd64",
					"9.12.1+1",
					"",
				),
			},
		),
	)
	modified := time.Now().Add(-time.Hour)
	o.NationalDatabase.SeedScore(
		vulnerabilityConstant.FixtureIdentifier,
		7.5,
		"HIGH",
		modified,
	)
	o.NationalDatabase.SeedScore("CVE-2099-20002", 9.8, "CRITICAL", modified)
	o.NationalDatabase.SeedScore("CVE-2099-20001", 9.0, "CRITICAL", modified)
	assert.Count(o.t, 0, o.Score().Sync(o.t.Context()))

	return o
}
