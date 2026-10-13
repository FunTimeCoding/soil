package relay

import (
	"encoding/binary"
	"github.com/funtimecoding/soil/pkg/tool/goffrelay/constant"
	"io"
)

func writeFrame(
	w io.Writer,
	b []byte,
) error {
	h := make([]byte, constant.HeaderLength)
	binary.NativeEndian.PutUint32(h, uint32(len(b)))

	if _, e := w.Write(h); e != nil {
		return e
	}

	_, e := w.Write(b)

	return e
}
