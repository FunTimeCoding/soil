package why_report

import "github.com/funtimecoding/soil/pkg/salt/types/chain"

type Report struct {
	State  string         `json:"state"`
	Status string         `json:"status"`
	Chains []*chain.Chain `json:"chains"`
}
