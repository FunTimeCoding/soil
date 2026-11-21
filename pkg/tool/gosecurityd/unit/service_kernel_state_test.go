package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/tool/gosecurityd/constant"
	"github.com/funtimecoding/soil/pkg/tool/gosecurityd/store/installed"
	"github.com/funtimecoding/soil/pkg/tool/gosecurityd/store/kernel_offer"
	"github.com/funtimecoding/soil/pkg/tool/gosecurityd/store/system"
	"github.com/funtimecoding/soil/pkg/tool/gosecurityd/unit/security_tester"
	"testing"
	"time"
)

func TestKernelStatePerHost(t *testing.T) {
	s := security_tester.New(t).KernelFleet().System()

	for name, state := range map[string]string{
		"kilo":     constant.KernelReboot,
		"lima":     constant.KernelUpgrade,
		"mike":     constant.KernelCurrent,
		"november": constant.KernelCurrent,
		"oscar":    constant.KernelStuck,
		"papa":     constant.KernelUnknown,
	} {
		d, e := s.Detail(name)
		assert.FatalOnError(t, e)
		assert.String(t, state, d.Summary.KernelState)
	}

	oscar, f := s.Detail("oscar")
	assert.FatalOnError(t, f)
	assert.String(t, "sources not reported yet", oscar.Summary.KernelReason)
	kilo, g := s.Detail("kilo")
	assert.FatalOnError(t, g)
	assert.String(t, "", kilo.Summary.KernelReason)
}

func TestStuckKernelNamesItsReason(t *testing.T) {
	s := security_tester.New(t).KernelFleet().System()
	stuck := func(
		name string,
		lists time.Time,
		origin string,
	) string {
		t.Helper()
		v := system.New(name, "bookworm", 0)
		v.Kernel = "9.12.38+deb12-amd64"
		v.ListsUpdated = &lists
		v.KernelOffers = []*kernel_offer.Offer{
			kernel_offer.New(
				"9.12.38-1~bpo12+1",
				100,
				[]string{"/var/lib/dpkg/status"},
			),
			kernel_offer.New("9.1.140-1", 500, []string{origin}),
		}
		assert.FatalOnError(
			t,
			s.Put(
				v,
				[]installed.Installed{
					*installed.New(
						"linux-image-9.12.38+deb12-amd64",
						"9.12.38-1~bpo12+1",
						"linux-signed-amd64",
						"9.12.38+1~bpo12+1",
						"",
					),
				},
			),
		)
		d, e := s.Detail(name)
		assert.FatalOnError(t, e)
		assert.String(t, "no upgrade path", d.Summary.KernelState)

		return d.Summary.KernelReason
	}
	fresh := time.Now().Add(-time.Hour)
	stale := time.Now().Add(-90 * 24 * time.Hour)
	assert.String(
		t,
		"no backports source",
		stuck("romeo", fresh, "http://deb.debian.org/debian bookworm/main"),
	)
	assert.String(
		t,
		"package lists 90 days old",
		stuck(
			"sierra",
			stale,
			"http://deb.debian.org/debian bookworm-backports/main",
		),
	)
	assert.String(
		t,
		"sources offer up to 9.12.38-1",
		stuck(
			"tango",
			fresh,
			"http://deb.debian.org/debian bookworm-backports/main",
		),
	)
}
