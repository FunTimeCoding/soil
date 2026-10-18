package unit

import (
	"context"
	"errors"
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/peer/retry"
	"testing"
	"time"
)

func TestSucceedsFirstAttempt(t *testing.T) {
	r, e := retry.Do(
		3,
		time.Millisecond,
		func() error {
			return nil
		},
	)
	assert.Nil(t, e)
	assert.Integer(t, 1, r.Attempts)
	assert.True(t, r.Elapsed >= 0)
}

func TestSucceedsAfterRetries(t *testing.T) {
	calls := 0
	r, e := retry.Do(
		3,
		time.Millisecond,
		func() error {
			calls++

			if calls < 3 {
				return errors.New("transient")
			}

			return nil
		},
	)
	assert.Nil(t, e)
	assert.Integer(t, 3, r.Attempts)
	assert.True(t, r.Elapsed >= 2*time.Millisecond)
}

func TestUntilReturnsOnceTheProbeSucceeds(t *testing.T) {
	calls := 0
	r, e := retry.Until(
		t.Context(),
		time.Second,
		time.Millisecond,
		func(context.Context) error {
			calls++

			if calls < 3 {
				return errors.New("unanswered")
			}

			return nil
		},
	)
	assert.Nil(t, e)
	assert.Integer(t, 3, r.Attempts)
	assert.True(t, r.Elapsed >= 2*time.Millisecond)
}

func TestUntilGivesUpAtTheDeadline(t *testing.T) {
	fail := errors.New("unanswered")
	r, e := retry.Until(
		t.Context(),
		10*time.Millisecond,
		time.Millisecond,
		func(context.Context) error {
			return fail
		},
	)
	assert.True(t, r.Attempts > 1)
	assert.True(t, r.Elapsed < 20*time.Millisecond)
	assert.StringContains(t, "after 10ms", e.Error())
	assert.True(t, errors.Is(e, fail))
}

func TestUntilStopsWhenItsContextIsCancelled(t *testing.T) {
	fail := errors.New("unanswered")
	x, cancel := context.WithCancel(t.Context())
	time.AfterFunc(10*time.Millisecond, cancel)
	start := time.Now()
	_, e := retry.Until(
		x,
		time.Minute,
		time.Second,
		func(context.Context) error {
			return fail
		},
	)
	assert.True(t, time.Since(start) < time.Second)
	assert.True(t, errors.Is(e, context.Canceled))
	assert.True(t, errors.Is(e, fail))
}

func TestUntilEndsAProbeThatHangsAtTheDeadline(t *testing.T) {
	start := time.Now()
	_, e := retry.Until(
		t.Context(),
		10*time.Millisecond,
		time.Millisecond,
		func(x context.Context) error {
			<-x.Done()

			return x.Err()
		},
	)
	assert.True(t, time.Since(start) < time.Second)
	assert.StringContains(t, "after 10ms", e.Error())
}

func TestExhausted(t *testing.T) {
	fail := errors.New("persistent")
	r, e := retry.Do(
		3,
		time.Millisecond,
		func() error {
			return fail
		},
	)
	assert.Integer(t, 3, r.Attempts)
	assert.True(t, r.Elapsed >= 2*time.Millisecond)
	assert.StringContains(t, "after 3 attempts", e.Error())
	assert.True(t, errors.Is(e, fail))
}
