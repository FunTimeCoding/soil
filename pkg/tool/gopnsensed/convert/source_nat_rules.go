package convert

import (
	"github.com/funtimecoding/soil/pkg/opnsense/source_nat"
	"github.com/funtimecoding/soil/pkg/tool/gopnsensed/generated/server"
)

func SourceNatRules(v []*source_nat.Rule) []server.SourceNatRule {
	result := []server.SourceNatRule{}

	for _, e := range v {
		result = append(result, *SourceNatRule(e))
	}

	return result
}
