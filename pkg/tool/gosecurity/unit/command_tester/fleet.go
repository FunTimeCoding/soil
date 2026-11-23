package command_tester

import (
	"github.com/funtimecoding/soil/pkg/tool/gosecurityd/client/mock_client"
	gosecuritydConstant "github.com/funtimecoding/soil/pkg/tool/gosecurityd/constant"
	"github.com/funtimecoding/soil/pkg/tool/gosecurityd/generated/client"
)

func (o *Tester) Fleet() *mock_client.Client {
	o.t.Helper()
	result := mock_client.New(o.t)

	for name, group := range map[string]string{"alfa": "red", "bravo": "blue"} {
		labels := map[string]string{"group": group}
		result.SeedSystemDetail(
			&client.SystemDetail{
				Summary: client.SystemSummary{
					Name:     name,
					Release:  "trixie",
					Coverage: gosecuritydConstant.CoverageSupported,
					TopScore: 9.8,
					Labels:   &labels,
				},
				Findings: []*client.SystemFinding{},
			},
		)
	}

	result.SeedUpgrade(
		client.UpgradeGroup{
			Release:         "trixie",
			Source:          "tango-utils",
			Hosts:           []client.UpgradeHost{{Name: "alfa"}},
			Vulnerabilities: []string{"CVE-2099-00001"},
			TopScore:        9.8,
		},
	)
	result.SeedUpgrade(
		client.UpgradeGroup{
			Release:  "trixie",
			Source:   "linux",
			Kernel:   true,
			Hosts:    []client.UpgradeHost{{Name: "bravo"}},
			TopScore: 7.5,
		},
	)
	result.SeedRemediationItem(
		client.RemediationItem{
			Subject:  "upgrade:trixie:tango-utils",
			Hosts:    []string{"alfa"},
			TopScore: 9.8,
		},
	)
	result.SeedRemediationItem(
		client.RemediationItem{
			Subject:  "CVE-2099-00002",
			Hosts:    []string{"bravo"},
			TopScore: 7.5,
		},
	)

	return result
}
