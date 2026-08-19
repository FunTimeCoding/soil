package armor

import (
	"encoding/pem"
	"github.com/funtimecoding/soil/pkg/errors/panics"
)

func decodeBlock(
	b []byte,
	kind string,
) []byte {
	block, _ := pem.Decode(b)

	if block == nil {
		panics.DecodeFail(kind)
	}

	if block.Type != kind {
		panics.Unexpected(block.Type)
	}

	return block.Bytes
}
