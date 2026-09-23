package unit

import (
	"github.com/funtimecoding/soil/pkg/alpine/constant"
	"github.com/funtimecoding/soil/pkg/alpine/index"
	"github.com/funtimecoding/soil/pkg/alpine/package_server"
	"github.com/funtimecoding/soil/pkg/alpine/unit/index_tester"
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/errors"
	"path/filepath"
	"testing"
)

func TestIndexRead(t *testing.T) {
	directory := t.TempDir()
	path := index_tester.WriteIndex(directory)
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
	index_tester.WriteIndex(
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
