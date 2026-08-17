package convert

import (
	"github.com/funtimecoding/soil/pkg/opnsense/rule"
	"github.com/funtimecoding/soil/pkg/tool/gopnsensed/generated/server"
)

func Rule(v *rule.Rule) *server.Rule {
	return &server.Rule{
		Identifier:      v.Identifier,
		Enabled:         v.Enabled,
		Sequence:        v.Sequence,
		Interface:       v.Interface,
		Direction:       v.Direction,
		Action:          v.Action,
		Protocol:        v.Protocol,
		SourceNet:       v.SourceNet,
		SourcePort:      v.SourcePort,
		DestinationNet:  v.DestinationNet,
		DestinationPort: v.DestinationPort,
		Log:             v.Log,
		Automatic:       v.Automatic,
		Description:     v.Description,
		Categories:      v.Categories,
	}
}
