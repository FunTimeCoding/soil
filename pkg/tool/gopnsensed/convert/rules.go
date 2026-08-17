package convert

import (
	"github.com/funtimecoding/soil/pkg/opnsense/rule"
	"github.com/funtimecoding/soil/pkg/tool/gopnsensed/generated/server"
)

func Rules(v []*rule.Rule) []server.Rule {
	result := []server.Rule{}

	for _, e := range v {
		result = append(result, *Rule(e))
	}

	return result
}
