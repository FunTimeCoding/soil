package signature

import (
	"bytes"
	"compress/gzip"
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/errors/validation"
	"github.com/funtimecoding/soil/pkg/system/constant"
	"io"
)

func FindGzipBoundary(payload []byte) (int, error) {
	for i := constant.AlpineGzipMinimumSize; i < len(payload)-constant.AlpineGzipHeaderSize; i++ {
		if payload[i] == 0x1f && payload[i+1] == 0x8b && payload[i+2] == 0x08 {
			gz, e := gzip.NewReader(bytes.NewReader(payload[i:]))

			if e != nil {
				continue
			}

			b := make([]byte, 1)
			_, e = gz.Read(b)
			errors.LogClose(gz)

			if e != nil && e != io.EOF {
				continue
			}

			return i, nil
		}
	}

	return 0, validation.New("not an apk package: no second gzip stream")
}
