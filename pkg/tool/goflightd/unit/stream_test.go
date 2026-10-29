package unit

import (
	"context"
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/system/run/stream/mock_stream"
	"github.com/funtimecoding/soil/pkg/tool/goflightd/constant"
	"github.com/funtimecoding/soil/pkg/tool/goflightd/unit/base"
	"testing"
	"time"
)

func TestTheStreamRecordsEachMessage(t *testing.T) {
	s := base.New(t)
	s.System.Serve(
		mock_stream.New(
			[]string{
				constant.SessionSample,
				"not json",
				constant.EmptyMessageSample,
				constant.SessionSample,
			},
			false,
			nil,
		),
	)
	assert.Count(t, 0, s.Service.CaptureStream(context.Background()))
	v, e := s.Service.Status()
	assert.FatalOnError(t, e)
	assert.Integer(t, 2, v.Events)
}

func TestAStreamThatEndsBadlyIsAnError(t *testing.T) {
	s := base.New(t)
	s.System.Serve(mock_stream.New(nil, false, errors.Format("exit", 1)))
	assert.Count(t, 1, s.Service.CaptureStream(context.Background()))
}

func TestCancelKillsTheStream(t *testing.T) {
	s := base.New(t)
	m := mock_stream.New(
		[]string{constant.SessionSample},
		true,
		errors.Format("killed", 9),
	)
	s.System.Serve(m)
	c, cancel := context.WithCancel(context.Background())
	done := make(chan []error)
	go func() { done <- s.Service.CaptureStream(c) }()
	cancel()

	select {
	case result := <-done:
		assert.Count(t, 0, result)
	case <-time.After(5 * time.Second):
		t.Fatal("capture outlived its cancel")
	}

	assert.Integer(t, 1, m.Killed())
}

func TestAFailingStoreStopsTheStream(t *testing.T) {
	s := base.New(t)
	m := mock_stream.New([]string{constant.SessionSample}, true, nil)
	s.System.Serve(m)
	s.Service.Close()
	done := make(chan []error)
	go func() { done <- s.Service.CaptureStream(context.Background()) }()

	select {
	case result := <-done:
		assert.Count(t, 1, result)
	case <-time.After(5 * time.Second):
		assert.FatalOnError(t, m.Kill())
		t.Fatal("capture kept reading past a failing store")
	}

	assert.Integer(t, 1, m.Killed())
}
