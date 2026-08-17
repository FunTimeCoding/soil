package source_nat

import "github.com/funtimecoding/soil/pkg/opnsense/response"

func New(v response.SourceNatRule) *Rule {
	return &Rule{
		Identifier:     v.Identifier,
		Enabled:        bool(v.Enabled),
		Interface:      v.Interface,
		Protocol:       v.Protocol,
		SourceNet:      v.SourceNet,
		DestinationNet: v.DestinationNet,
		Target:         v.Target,
		TargetPort:     v.TargetPort,
		Log:            bool(v.Log),
		Description:    v.Description,
	}
}
