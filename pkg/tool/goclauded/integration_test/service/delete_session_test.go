package service

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/integration_test/service_tester"
	"os"
	"path/filepath"
	"testing"
)

func TestDeleteSessionClearsTrackerState(t *testing.T) {
	s := service_tester.New(t)
	writeSessionFile(s.Harbor, "doomed", "some-slug")
	s.Service.PopulateCache()
	s.Service.CheckConsistency()
	assert.True(t, s.Store.GetSession("doomed") != nil)
	errors.PanicOnError(s.Service.DeleteSession("doomed"))
	// the mock claude client does not remove the harbor file
	errors.PanicOnError(os.Remove(filepath.Join(s.Harbor, "doomed.jsonl")))
	_, tracked := s.Store.Store.TrackerStates()["doomed"]
	assert.False(t, tracked)
	s.Service.PopulateCache()
	s.Service.CheckConsistency()
	assert.True(t, s.Store.GetSession("doomed") == nil)
}
