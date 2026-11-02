package unit

import (
	"errors"
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/gw2/log_manager/log"
	"github.com/funtimecoding/soil/pkg/gw2/log_manager/response"
	"github.com/funtimecoding/soil/pkg/relational/lite"
	"github.com/funtimecoding/soil/pkg/tool/goraidd/store"
	"github.com/funtimecoding/soil/pkg/tool/goraidd/unit/service_tester"
	"io/fs"
	"testing"
)

func TestSyncingTheLogCacheRecordsEachFightOnce(t *testing.T) {
	cache := t.TempDir()
	service_tester.WriteLogCache(
		cache,
		"20261024-201500.zevtc",
		"20261024-203000.zevtc",
	)
	s := service_tester.New(t, cache, t.TempDir())
	s.SyncLogCache()
	s.SyncLogCache()
	assert.Count(t, 2, s.Fights())
}

func TestARecordedFightIsNotCountedAsNewAgain(t *testing.T) {
	s := store.New(lite.NewMemory())
	logs := log.NewSlice(
		[]*response.Log{{FileName: "20261024-201500.zevtc", MapIdentifier: 38}},
	)
	assert.Integer(t, 1, s.SaveLogEntries(logs))
	assert.Integer(t, 0, s.SaveLogEntries(logs))
}

func TestTheServiceStartsWithoutALogCache(t *testing.T) {
	s := service_tester.New(t, t.TempDir(), t.TempDir())
	assert.Count(t, 0, s.Fights())
	failures := s.SyncLogCache()
	assert.Count(t, 1, failures)
	assert.True(t, errors.Is(failures[0], fs.ErrNotExist))
}
