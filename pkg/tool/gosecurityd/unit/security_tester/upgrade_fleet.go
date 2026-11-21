package security_tester

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/tool/gosecurityd/constant"
	"github.com/funtimecoding/soil/pkg/tool/gosecurityd/store/installed"
	"github.com/funtimecoding/soil/pkg/tool/gosecurityd/store/system"
	vulnerabilityConstant "github.com/funtimecoding/soil/pkg/vulnerability/constant"
)

func (o *Tester) UpgradeFleet() *Tester {
	o.t.Helper()

	for _, v := range [][4]string{
		{
			vulnerabilityConstant.FixtureIdentifier,
			"tango-utils",
			vulnerabilityConstant.DebianResolved,
			"8.6.1-1",
		},
		{
			vulnerabilityConstant.FixtureIdentifier,
			"romeo",
			vulnerabilityConstant.DebianOpen,
			"",
		},
		{
			"CVE-2099-20002",
			"tango-utils",
			vulnerabilityConstant.DebianResolved,
			"8.6.0-1",
		},
		{
			"CVE-2099-20001",
			constant.KernelSource,
			vulnerabilityConstant.DebianResolved,
			"9.12.2-1",
		},
		{
			"CVE-2099-20003",
			"linux-signed-amd64",
			vulnerabilityConstant.DebianResolved,
			"9.12.2+1",
		},
	} {
		o.SecurityTracker.SeedStatus(v[0], v[1], "trixie", v[2], v[3])
	}

	o.ExtendedTracker.SeedStatus(
		vulnerabilityConstant.FixtureIdentifier,
		"tango-utils",
		"bullseye",
		vulnerabilityConstant.DebianResolved,
		"8.2.5-2.1~deb11u9",
	)
	assert.Count(o.t, 0, o.Debian().Refresh(o.t.Context()))
	s := o.System()
	assert.FatalOnError(
		o.t,
		s.Put(
			system.New("alfa", "trixie", 2),
			[]installed.Installed{
				*installed.New(
					"libtango5",
					"8.4.0-1",
					"tango-utils",
					"8.4.0-1",
					"8.6.1-1",
				),
				*installed.New(
					"tango-utils",
					"8.4.0-1",
					"tango-utils",
					"8.4.0-1",
					"8.6.1-1",
				),
				*installed.New(
					"libtango-dev",
					"8.4.0-1",
					"tango-utils",
					"8.4.0-1",
					"",
				),
				*installed.New("romeo", "1.0-1", "romeo", "1.0-1", ""),
				*installed.New(
					"linux-libc-dev",
					"9.12.1-1",
					constant.KernelSource,
					"9.12.1-1",
					"9.12.2-1",
				),
			},
		),
	)
	assert.FatalOnError(
		o.t,
		s.Put(
			system.New("bravo", "trixie", 1),
			[]installed.Installed{
				*installed.New(
					"libtango5",
					"8.2.0-1",
					"tango-utils",
					"8.2.0-1",
					"8.6.1-1",
				),
				*installed.New(
					"linux-image-9.12.1-amd64",
					"9.12.1+1",
					"linux-signed-amd64",
					"9.12.1+1",
					"9.12.2+1",
				),
			},
		),
	)
	assert.FatalOnError(
		o.t,
		s.Put(
			system.New("charlie", "bullseye", 1),
			[]installed.Installed{
				*installed.New(
					"libtango5",
					"8.2.5-2.1~deb11u1",
					"tango-utils",
					"8.2.5-2.1~deb11u1",
					"8.2.5-2.1~deb11u9",
				),
			},
		),
	)
	assert.FatalOnError(
		o.t,
		s.Put(
			system.New("delta", "trixie", 0),
			[]installed.Installed{
				*installed.New(
					"libtango5",
					"8.4.0-1",
					"tango-utils",
					"8.4.0-1",
					"",
				),
			},
		),
	)

	return o
}
