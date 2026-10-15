package unit

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/identity"
	"github.com/funtimecoding/soil/pkg/instrument"
	"github.com/funtimecoding/soil/pkg/instrument/unit/sentry_receiver"
	stringsConstant "github.com/funtimecoding/soil/pkg/strings/constant"
	"testing"
	"time"
)

func TestFlushDeliversReportedEvent(t *testing.T) {
	received := sentry_receiver.New(t)
	i := instrument.New(
		identity.New(
			stringsConstant.LowerAlfa,
			stringsConstant.LowerAlfa,
			stringsConstant.LowerAlfa,
		),
	)
	l := captureLog(t)
	func() {
		defer func() { i.Flush(recover()) }()
		i.Reporter().CaptureException(fmt.Errorf("alfa"))
	}()
	assert.Integer(t, 1, int(received.Load()))
	assert.StringNotContains(t, "sentry flush fail", l.String())
}

func TestCloseDeliversEventsQueuedBehindASlowSend(t *testing.T) {
	received := sentry_receiver.NewDelayed(t, 20*time.Millisecond)
	i := instrument.New(
		identity.New(
			stringsConstant.LowerAlfa,
			stringsConstant.LowerAlfa,
			stringsConstant.LowerAlfa,
		),
	)

	for range 10 {
		i.Reporter().CaptureException(fmt.Errorf("alfa"))
	}

	i.Close()
	assert.Integer(t, 10, int(received.Load()))
}
