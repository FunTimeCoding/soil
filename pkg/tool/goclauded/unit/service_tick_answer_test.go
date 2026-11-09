package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/unit/service_tester"
	"testing"
)

func TestDeliveryWithDatabaseGoneAnswersFailure(t *testing.T) {
	s := service_tester.New(t)
	s.Close()
	assert.Count(t, 1, s.Memory.DeliverImpressions())
}

func TestCruiseReleaseWithDatabaseGoneAnswersFailure(t *testing.T) {
	s := service_tester.New(t)
	s.Close()
	assert.Count(t, 1, s.Message.ReleaseCruise())
}
