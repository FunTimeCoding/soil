package why_report

import "github.com/funtimecoding/soil/pkg/salt/types/chain"

func New(
	state string,
	status string,
	chains []*chain.Chain,
) *Report {
	return &Report{State: state, Status: status, Chains: chains}
}
