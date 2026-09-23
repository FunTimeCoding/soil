package service

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/integration/service_tester"
	"testing"
)

func TestDeleteSessionClearsTrackerState(t *testing.T) {
	s := service_tester.New(t)
	s.WriteSessionFile("doomed", "some-slug")
	s.Service.PopulateCache()
	s.Service.CheckConsistency()
	assert.True(t, s.Store.GetSession("doomed") != nil)
	_, e := s.Service.DeleteSession("doomed", s.Service.DeleteHash("doomed"))
	errors.PanicOnError(e)
	_, tracked := s.Store.Store.TrackerStates()["doomed"]
	assert.False(t, tracked)
	s.Service.PopulateCache()
	s.Service.CheckConsistency()
	assert.True(t, s.Store.GetSession("doomed") == nil)
}
