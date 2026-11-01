package user

import "example/pkg/source"

var Common = []string{"alfa"}

func All() []string {
	return append(Common, source.AdditionalEntriesForTheVariadicSpread...)
}
