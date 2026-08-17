package convert

import (
	"github.com/funtimecoding/soil/pkg/opnsense/source_nat"
	"github.com/funtimecoding/soil/pkg/tool/gopnsensed/generated/server"
)

func SourceNatRule(v *source_nat.Rule) *server.SourceNatRule {
	return &server.SourceNatRule{
		Identifier:     v.Identifier,
		Enabled:        v.Enabled,
		Interface:      v.Interface,
		Protocol:       v.Protocol,
		SourceNet:      v.SourceNet,
		DestinationNet: v.DestinationNet,
		Target:         v.Target,
		TargetPort:     v.TargetPort,
		Log:            v.Log,
		Description:    v.Description,
	}
}
