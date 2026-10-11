package signature

import (
	"bytes"
	"compress/gzip"
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/system/constant"
	"io"
	"log"
)

func FindGzipBoundary(payload []byte) int {
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

			return i
		}
	}

	log.Panicf("could not find second gzip stream")

	return 0
}
