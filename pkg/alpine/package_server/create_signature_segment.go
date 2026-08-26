package package_server

import (
	"archive/tar"
	"bytes"
	"compress/gzip"
	"github.com/funtimecoding/soil/pkg/alpine/constant"
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/strings/join"
	"github.com/funtimecoding/soil/pkg/system"
)

func CreateSignatureSegment(
	signature []byte,
	keyName string,
) []byte {
	var b bytes.Buffer
	z := gzip.NewWriter(&b)
	t := tar.NewWriter(z)
	system.TarWriteHeader(
		t,
		&tar.Header{
			Name: join.Empty(
				constant.SignaturePrefix,
				keyName,
				constant.PublicKeySuffix,
			),
			Size: int64(len(signature)),
			Mode: 0644,
			Uid:  0,
			Gid:  0,
		},
	)
	system.TarWrite(t, signature)
	errors.PanicFlush(t)
	errors.PanicClose(z)

	return b.Bytes()
}
