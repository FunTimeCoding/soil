package unit

import (
	"archive/tar"
	"bytes"
	"compress/gzip"
	"github.com/funtimecoding/soil/pkg/alpine/constant"
	"github.com/funtimecoding/soil/pkg/alpine/index"
	"github.com/funtimecoding/soil/pkg/alpine/package_server"
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/system"
	"path/filepath"
	"testing"
)

func TestIndexRead(t *testing.T) {
	directory := t.TempDir()
	path := writeIndex(directory)
	entries, e := index.Read(path)
	errors.PanicOnError(e)
	assert.Count(t, 2, entries)
	assert.String(t, "gohw", entries[0].Name)
	assert.String(t, "0.11.96-r1", entries[0].Version)
	assert.String(t, "x86_64", entries[0].Architecture)
	assert.String(t, "gobuild", entries[1].Name)
}

func TestIndexes(t *testing.T) {
	directory := t.TempDir()
	writeIndex(
		filepath.Join(directory, "rolling", "main", constant.Architecture),
	)
	listings, e := package_server.Indexes(directory)
	errors.PanicOnError(e)
	assert.Count(t, 1, listings)
	assert.String(t, "rolling", listings[0].Version)
	assert.String(t, "main", listings[0].Repository)
	assert.String(t, "x86_64", listings[0].Architecture)
	assert.Count(t, 2, listings[0].Packages)
}

func writeIndex(directory string) string {
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
