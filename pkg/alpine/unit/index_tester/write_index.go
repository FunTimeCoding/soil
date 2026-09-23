package index_tester

import (
	"archive/tar"
	"bytes"
	"compress/gzip"
	"github.com/funtimecoding/soil/pkg/alpine/constant"
	"github.com/funtimecoding/soil/pkg/alpine/package_server"
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/system"
	"path/filepath"
)

func WriteIndex(directory string) string {
	system.MakeDirectory(directory)
	path := filepath.Join(directory, constant.IndexArchive)
	content := "C:Q1checksum\nP:gohw\nV:0.11.96-r1\nA:x86_64\n\nP:gobuild\nV:0.11.95-r1\nA:x86_64\n"
	var b bytes.Buffer
	b.Write(
		package_server.CreateSignatureSegment([]byte("signature"), "test.rsa"),
	)
	z := gzip.NewWriter(&b)
	w := tar.NewWriter(z)
	system.TarWriteHeader(
		w,
		&tar.Header{
			Name: constant.IndexFile,
			Size: int64(len(content)),
			Mode: 0644,
		},
	)
	system.TarWrite(w, []byte(content))
	errors.PanicClose(w)
	errors.PanicClose(z)
	system.WriteFile(path, b.Bytes(), 0644)

	return path
}
