package security_tester

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/tool/gosecurityd/constant"
	"github.com/funtimecoding/soil/pkg/tool/gosecurityd/store/installed"
	"github.com/funtimecoding/soil/pkg/tool/gosecurityd/store/kernel_offer"
	"github.com/funtimecoding/soil/pkg/tool/gosecurityd/store/system"
	vulnerabilityConstant "github.com/funtimecoding/soil/pkg/vulnerability/constant"
	"time"
)

func (o *Tester) RemediationFleet() *Tester {
	o.t.Helper()
	o.SecurityTracker.SeedStatus(
		"CVE-2099-20005",
		"tango-utils",
		"trixie",
		vulnerabilityConstant.DebianOpen,
		"",
	)
	o.SecurityTracker.SeedUrgency(
		"CVE-2099-20005",
		"tango-utils",
		"trixie",
		constant.Unimportant,
	)
	o.SecurityTracker.SeedStatus(
		"CVE-2099-20006",
		constant.KernelSource,
		"trixie",
		vulnerabilityConstant.DebianOpen,
		"",
	)
	o.SecurityTracker.SeedStatus(
		"CVE-2099-20007",
		constant.KernelSource,
		"trixie",
		vulnerabilityConstant.DebianResolved,
		"9.12.9-1",
	)

	for identifier, score := range map[string]float64{
		"CVE-2099-20005": 9.5,
		"CVE-2099-20006": 9.4,
		"CVE-2099-20007": 9.2,
	} {
		o.NationalDatabase.SeedScore(
			identifier,
			score,
			"CRITICAL",
			time.Now().Add(-time.Hour),
		)
	}

	o.ScoredFleet()
	o.putKernel(
		"foxtrot",
		"trixie",
		"9.12.1-amd64",
		*installed.New(
			"linux-image-9.12.1-amd64",
			"9.12.1-1",
			"linux-signed-amd64",
			"9.12.1+1",
			"",
		),
		*installed.New(
			"linux-image-9.12.2-amd64",
			"9.12.2-1",
			"linux-signed-amd64",
			"9.12.2+1",
			"",
		),
	)
	hotel := system.New("hotel", "trixie", 0)
	hotel.Kernel = "9.12.1-amd64"
	hotel.KernelOffers = []*kernel_offer.Offer{
		kernel_offer.New("9.12.2-1", 500, []string{"trixie-security"}),
	}
	assert.FatalOnError(
		o.t,
		o.System().Put(
			hotel,
			[]installed.Installed{
				*installed.New(
					"linux-image-9.12.1-amd64",
					"9.12.1-1",
					"linux-signed-amd64",
					"9.12.1+1",
					"",
				),
				*installed.New(
					"linux-image-amd64",
					"9.12.1-1",
					"linux-signed-amd64",
					"9.12.1+1",
					"9.12.2-1",
				),
			},
		),
	)
	assert.FatalOnError(o.t, o.System().Put(system.New("golf", "kilo", 0), nil))
	assert.FatalOnError(
		o.t,
		o.System().Put(
			system.New("delta", "trixie", 0),
			[]installed.Installed{
				*installed.New(
					"libtango5",
					"8.4.0-1",
					"tango-utils",
					"8.4.0-1",
					"",
				),
				*installed.New("romeo", "1.0-1", "romeo", "1.0-1", ""),
			},
		),
	)

	return o
}
