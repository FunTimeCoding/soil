package security_tester

import (
	vulnerabilityConstant "github.com/funtimecoding/soil/pkg/vulnerability/constant"
	"time"
)

func (o *Tester) SeedScores() *Tester {
	o.t.Helper()

	for i, v := range []float64{9.8, 7.5, 5.0, 3.1, 10.0} {
		o.NationalDatabase.SeedScore(
			[]string{
				"CVE-2099-90001",
				"CVE-2099-90002",
				"CVE-2099-90003",
				"CVE-2099-90004",
				vulnerabilityConstant.FixtureIdentifier,
			}[i],
			v,
			"HIGH",
			time.Now().Add(-time.Duration(i+1)*24*time.Hour),
		)
	}

	return o
}
