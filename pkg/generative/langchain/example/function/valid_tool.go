package function

import (
	"github.com/funtimecoding/soil/pkg/generative/constant"
	"slices"
)

func validTool(name string) bool {
	var valid []string

	for _, v := range constant.LangchainExampleFunctions {
		valid = append(valid, v.Name)
	}

	return slices.Contains(valid, name)
}
