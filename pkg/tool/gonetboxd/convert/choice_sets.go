package convert

import (
	"github.com/funtimecoding/soil/pkg/netbox/choice_set"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/generated/server"
)

func ChoiceSets(v []*choice_set.Set) []*server.ChoiceSet {
	result := make([]*server.ChoiceSet, 0, len(v))

	for _, s := range v {
		result = append(result, ChoiceSet(s))
	}

	return result
}
