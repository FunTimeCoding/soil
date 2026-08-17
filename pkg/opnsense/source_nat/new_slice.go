package source_nat

import "github.com/funtimecoding/soil/pkg/opnsense/response"

func NewSlice(v []response.SourceNatRule) []*Rule {
	var result []*Rule

	for _, e := range v {
		result = append(result, New(e))
	}

	return result
}
