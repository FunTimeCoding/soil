package signature

import (
	"archive/tar"
	"bytes"
	"compress/gzip"
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/strings/join"
	"github.com/funtimecoding/soil/pkg/system"
	"github.com/funtimecoding/soil/pkg/system/constant"
)

func CreateSignatureSegment(
	signature []byte,
	keyName string,
) []byte {
	segment := system.TarSegment(
		func(t *tar.Writer) {
			system.TarWriteHeader(
				t,
				&tar.Header{
					Name: join.Empty(
						constant.AlpineSignaturePrefix,
						keyName,
						constant.AlpinePublicKeySuffix,
					),
					Size: int64(len(signature)),
					Mode: 0644,
					Uid:  0,
					Gid:  0,
				},
			)
			system.TarWrite(t, signature)
		},
	)
	var b bytes.Buffer
	z := gzip.NewWriter(&b)
	system.Copy(bytes.NewReader(segment), z)
	errors.PanicClose(z)

	return b.Bytes()
}
