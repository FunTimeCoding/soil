package relay

import (
	"encoding/binary"
	"fmt"
	"github.com/funtimecoding/soil/pkg/tool/goffrelay/constant"
	"io"
)

func readFrame(r io.Reader) ([]byte, error) {
	h := make([]byte, constant.HeaderLength)

	if _, e := io.ReadFull(r, h); e != nil {
		return nil, e
	}

	n := binary.NativeEndian.Uint32(h)

	if n > constant.FrameLimit {
		return nil, fmt.Errorf("frame of %d bytes exceeds the limit", n)
	}

	b := make([]byte, n)

	if _, e := io.ReadFull(r, b); e != nil {
		return nil, e
	}

	return b, nil
}
