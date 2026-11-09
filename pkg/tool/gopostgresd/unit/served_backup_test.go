package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/notation"
	"github.com/funtimecoding/soil/pkg/postgres/types/result"
	"github.com/funtimecoding/soil/pkg/tool/gopostgresd/constant"
	"github.com/funtimecoding/soil/pkg/tool/gopostgresd/generated/client"
	"net/http"
	"testing"
)

func TestRestDeleteOfAbsentBackupIsNotFound(t *testing.T) {
	_, c, r, _ := served(t, nil)
	v, e := c.DeleteBackupWithResponse(
		t.Context(),
		"2030-01-02T03-04-05",
		&client.DeleteBackupParams{Instance: new("alfa")},
	)
	assert.FatalOnError(t, e)
	assert.Integer(t, http.StatusNotFound, v.StatusCode())
	assert.Count(t, 0, r.Events())
}

func TestRestBackupNameOutsidePatternIsRefused(t *testing.T) {
	_, c, r, _ := served(t, nil)
	v, e := c.RestoreBackupWithResponse(
		t.Context(),
		"../alfa",
		client.RestoreBackupJSONRequestBody{Instance: new("alfa")},
	)
	assert.FatalOnError(t, e)
	assert.Integer(t, http.StatusBadRequest, v.StatusCode())
	assert.Count(t, 0, r.Events())
}

func TestRestRestoreInPlaceWhileOthersAreConnectedIsConflict(t *testing.T) {
	_, c, r, postgres := served(t, nil)
	created, e := c.CreateBackupWithResponse(
		t.Context(),
		client.CreateBackupJSONRequestBody{Instance: new("alfa")},
	)
	assert.FatalOnError(t, e)
	assert.Integer(t, http.StatusOK, created.StatusCode())
	postgres.SeedResult(
		constant.ConnectedSessionsQuery,
		result.New(
			nil,
			[][]any{{"bravo", "alfa", "10.0.0.5", "2030-01-02T00:00:00Z"}},
		),
	)
	v, f := c.RestoreBackupWithResponse(
		t.Context(),
		created.JSON200.Name,
		client.RestoreBackupJSONRequestBody{
			Instance: new("alfa"),
			Target:   new("alfa"),
		},
	)
	assert.FatalOnError(t, f)
	assert.Integer(t, http.StatusConflict, v.StatusCode())
	assert.StringContains(t, "bravo (alfa@10.0.0.5", v.JSON409.Error)
	assert.Count(t, 0, r.Events())
}

func TestModelContextBackupIsCreatedListedAndDeleted(t *testing.T) {
	s, _, r, _ := served(t, nil)
	m := s.Connect(t)
	defer m.Close()
	var created client.Backup
	notation.MustDecodeBytes(
		[]byte(m.MustCallTool(constant.CreateBackup, nil)),
		&created,
		false,
	)
	assert.String(t, "alfa", created.Instance)
	var listed client.BackupList
	notation.MustDecodeBytes(
		[]byte(m.MustCallTool(constant.ListBackups, nil)),
		&listed,
		false,
	)
	assert.Count(t, 1, listed.Backups)
	assert.String(t, created.Name, listed.Backups[0].Name)
	m.MustCallTool(
		constant.DeleteBackup,
		map[string]any{"backup": created.Name},
	)
	notation.MustDecodeBytes(
		[]byte(m.MustCallTool(constant.ListBackups, nil)),
		&listed,
		false,
	)
	assert.Count(t, 0, listed.Backups)
	assert.Count(t, 0, r.Events())
}

func TestModelContextRestoreOfAbsentBackupIsQuiet(t *testing.T) {
	s, _, r, _ := served(t, nil)
	m := s.Connect(t)
	defer m.Close()
	m.MustCallToolError(
		constant.RestoreBackup,
		map[string]any{"backup": "2030-01-02T03-04-05"},
	)
	assert.Count(t, 0, r.Events())
}
