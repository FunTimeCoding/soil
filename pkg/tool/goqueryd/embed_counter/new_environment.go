package embed_counter

import (
	"github.com/funtimecoding/soil/pkg/system/environment"
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/constant"
)

func NewEnvironment() (*Counter, error) {
	return New(environment.Required(constant.EmbedTokenizerEnvironment))
}
