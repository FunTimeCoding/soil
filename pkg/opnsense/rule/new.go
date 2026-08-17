package rule

import "github.com/funtimecoding/soil/pkg/opnsense/response"

func New(v response.Rule) *Rule {
	return &Rule{
		Identifier:      v.Identifier,
		Enabled:         bool(v.Enabled),
		Sequence:        v.Sequence,
		Interface:       v.Interface,
		Direction:       v.Direction,
		Action:          v.Action,
		Protocol:        v.Protocol,
		SourceNet:       v.SourceNet,
		SourcePort:      v.SourcePort,
		DestinationNet:  v.DestinationNet,
		DestinationPort: v.DestinationPort,
		Log:             bool(v.Log),
		Automatic:       bool(v.Automatic),
		Description:     v.Description,
		Categories:      v.Categories,
	}
}
