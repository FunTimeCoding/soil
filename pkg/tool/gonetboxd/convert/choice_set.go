package convert

import (
	"github.com/funtimecoding/soil/pkg/netbox/choice_set"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/generated/server"
)

func ChoiceSet(v *choice_set.Set) *server.ChoiceSet {
	return &server.ChoiceSet{
		Identifier: v.Identifier,
		Name:       v.Name,
		Choices:    v.Choices,
	}
}
