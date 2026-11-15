//go:build local

package integration

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/strings/join"
	"github.com/funtimecoding/soil/pkg/tool/gopostgresd/generated/client"
	"github.com/funtimecoding/soil/pkg/tool/gopostgresd/integration/postgres_tester"
	"net/http"
	"testing"
)

func TestBackupRestoresBesideAndInPlace(t *testing.T) {
	directory := t.TempDir()
	container, i := postgres_tester.Start(t, directory)
	postgres_tester.Psql(
		t,
		container,
		i,
		"alfa",
		"CREATE TABLE item (name text); INSERT INTO item VALUES ('alfa'), ('bravo'), ('charlie')",
	)
	c, r := postgres_tester.Serve(t, container, i, directory)
	created, e := c.CreateBackupWithResponse(
		t.Context(),
		client.CreateBackupJSONRequestBody{},
	)
	assert.FatalOnError(t, e)
	assert.Integer(t, http.StatusOK, created.StatusCode())
	assert.Integer(t, 1, *created.JSON200.Tables)
	name := created.JSON200.Name
	listed, f := c.ListBackupsWithResponse(
		t.Context(),
		&client.ListBackupsParams{},
	)
	assert.FatalOnError(t, f)
	assert.Count(t, 1, listed.JSON200.Backups)
	beside, g := c.RestoreBackupWithResponse(
		t.Context(),
		name,
		client.RestoreBackupJSONRequestBody{},
	)
	assert.FatalOnError(t, g)
	assert.Integer(t, http.StatusOK, beside.StatusCode())
	assert.String(t, join.Empty("alfa_restore_", name), beside.JSON200.Database)
	assert.String(
		t,
		"3",
		postgres_tester.Psql(
			t,
			container,
			i,
			beside.JSON200.Database,
			"SELECT count(*) FROM item",
		),
	)
	postgres_tester.Psql(t, container, i, "alfa", "DELETE FROM item")
	place, h := c.RestoreBackupWithResponse(
		t.Context(),
		name,
		client.RestoreBackupJSONRequestBody{Target: new("alfa")},
	)
	assert.FatalOnError(t, h)
	assert.Integer(t, http.StatusOK, place.StatusCode())
	assert.String(
		t,
		"3",
		postgres_tester.Psql(
			t,
			container,
			i,
			"alfa",
			"SELECT count(*) FROM item",
		),
	)
	removed, j := c.DeleteBackupWithResponse(
		t.Context(),
		name,
		&client.DeleteBackupParams{},
	)
	assert.FatalOnError(t, j)
	assert.Integer(t, http.StatusNoContent, removed.StatusCode())
	assert.Count(t, 0, r.Events())
}
