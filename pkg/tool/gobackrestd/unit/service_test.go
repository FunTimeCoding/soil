package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/backrest/mock_client"
	"github.com/funtimecoding/soil/pkg/tool/gobackrestd/unit/service_tester"
	"testing"
)

func TestLatestKeepsNewestBackupPerType(t *testing.T) {
	s := fixtureStanzas("backed_up.json")[0]
	older := fixtureStanzas("backed_up.json")[0].Backup[0]
	older.Label = "older"
	older.Timestamp.Stop = 1
	s.Backup = append(s.Backup, older)
	c := mock_client.New()
	c.Add(s)
	v, e := service_tester.New(t, c).Latest(t.Context())
	assert.FatalOnError(t, e)
	assert.Count(t, 1, v)
	assert.Integer(t, 0, v[0].Status)
	assert.Integer(t, 2, len(v[0].Latest))
	assert.String(t, "20261111-010538F", v[0].Latest["full"].Label)
	assert.String(
		t,
		"20261111-010538F_20261111-010540D",
		v[0].Latest["diff"].Label,
	)
}

func TestLatestCarriesStatusOfStanzaWithoutBackups(t *testing.T) {
	c := mock_client.New()
	c.Add(fixtureStanzas("no_backup.json")...)
	v, e := service_tester.New(t, c).Latest(t.Context())
	assert.FatalOnError(t, e)
	assert.Count(t, 1, v)
	assert.Integer(t, 2, v[0].Status)
	assert.Integer(t, 0, len(v[0].Latest))
}
