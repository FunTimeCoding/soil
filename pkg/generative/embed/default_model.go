package embed

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/generative/constant"
)

func DefaultModel(scheme string) string {
	switch scheme {
	case constant.OllamaScheme:
		return constant.OllamaEmbedModel
	case constant.OpenAIScheme:
		return constant.OpenAIEmbedModel
	}

	panic(fmt.Sprintf("unknown embed target scheme: %s", scheme))
}
