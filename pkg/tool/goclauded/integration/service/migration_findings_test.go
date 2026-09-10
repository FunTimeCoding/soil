package service

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/constant"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/integration/service_tester"
	"testing"
)

func TestFindingsReportUnkeyedRows(t *testing.T) {
	s := service_tester.New(t)
	assert.FatalOnError(
		t,
		s.Store.Store.Database().Exec(
			`INSERT INTO queue (callsign, kind, body, consumed, created_at)
			VALUES ('Wren', 'message', 'unkeyed', 0, '2026-09-02 10:00:00+00:00')`,
		).Error,
	)
	result := findingsByKind(t, s, constant.MigrationIncomplete)
	assert.Count(t, 1, result)
	assert.String(t, "queue", result[0].Subject)
	assert.Integer(t, 1, result[0].Count)
	assert.StringContains(t, "backfill", result[0].Detail)
}

func TestFindingsAreSilentWhenEveryRowCarriesItsSession(t *testing.T) {
	s := service_tester.New(t)
	r := s.Check("session-1")
	assert.FatalOnError(t, s.Service.Send(r.Callsign, r.Callsign, "hello"))
	assert.Count(t, 0, findingsByKind(t, s, constant.MigrationIncomplete))
}
