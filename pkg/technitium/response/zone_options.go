package response

import "github.com/funtimecoding/soil/pkg/technitium/policy"

type ZoneOptions struct {
	Name                 string           `json:"name"`
	Update               string           `json:"update"`
	UpdateNetwork        []string         `json:"updateNetworkACL"`
	UpdatePolicies       []*policy.Policy `json:"updateSecurityPolicies"`
	ZoneTransfer         string           `json:"zoneTransfer"`
	ZoneTransferNetwork  []string         `json:"zoneTransferNetworkACL"`
	ZoneTransferKeyNames []string         `json:"zoneTransferTsigKeyNames"`
	AvailableKeyNames    []string         `json:"availableTsigKeyNames"`
}
