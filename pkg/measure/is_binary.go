package measure

import (
	"bytes"
	"github.com/funtimecoding/soil/pkg/measure/constant"
)

func isBinary(b []byte) bool {
	if len(b) > constant.BinaryProbe {
		b = b[:constant.BinaryProbe]
	}

	return bytes.IndexByte(b, 0) >= 0
}
