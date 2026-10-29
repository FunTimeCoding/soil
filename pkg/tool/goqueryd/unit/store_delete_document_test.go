package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/relational/lite/connection"
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/constant"
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/store"
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/unit/store_tester"
	"testing"
)

func TestDeletingAnAbsentDocumentIsNoError(t *testing.T) {
	deleted, e := store_tester.OpenTestStore(t).DeleteDocument(
		"notes",
		"alfa.md",
	)
	assert.FatalOnError(t, e)
	assert.False(t, deleted)
}

func TestADeleteOnABrokenDatabaseFails(t *testing.T) {
	d := connection.NewMemory()
	s := store.New(d, constant.FixtureEmbedModel)
	errors.PanicClose(d)
	_, e := s.DeleteDocument("notes", "alfa.md")
	assert.Error(t, e)
}
