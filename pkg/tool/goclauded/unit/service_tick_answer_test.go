package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/unit/service_tester"
	"testing"
)

func TestADeliveryWithTheDatabaseGoneAnswersTheFailure(t *testing.T) {
	s := service_tester.New(t)
	s.Close()
	assert.Count(t, 1, s.Memory.DeliverImpressions())
}

func TestACruiseReleaseWithTheDatabaseGoneAnswersTheFailure(t *testing.T) {
	s := service_tester.New(t)
	s.Close()
	assert.Count(t, 1, s.Message.ReleaseCruise())
}
