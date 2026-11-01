package embed_counter

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/system/environment"
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/constant"
)

func NewEnvironment() (*Counter, error) {
	path := environment.Optional(constant.EmbedTokenizerEnvironment)

	if path == "" {
		return nil, fmt.Errorf(
			constant.EmbedTokenizerUnset,
			constant.EmbedTokenizerEnvironment,
			constant.EmbedTokenizerDefault,
			constant.EmbedTokenizerSource,
		)
	}

	result, e := New(path)

	if e != nil {
		return nil, fmt.Errorf(
			constant.EmbedTokenizerUnreadable,
			constant.EmbedTokenizerEnvironment,
			e,
			path,
			constant.EmbedTokenizerSource,
		)
	}

	return result, nil
}
