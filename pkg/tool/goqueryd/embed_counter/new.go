package embed_counter

import (
	"fmt"
	"github.com/amikos-tech/pure-tokenizers"
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/tokenizer"
)

func New(path string) (*Counter, error) {
	t, e := tokenizer.NewCounter(path)

	if e != nil {
		return nil, fmt.Errorf("load embed tokenizer: %w", e)
	}

	empty, f := t.Encode("", tokenizers.WithAddSpecialTokens())

	if f != nil {
		errors.PanicOnError(t.Close())

		return nil, fmt.Errorf("measure embed specials: %w", f)
	}

	return &Counter{tokenizer: t, specials: len(empty.IDs)}, nil
}
