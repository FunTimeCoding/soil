package unit

import (
	"errors"
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/errors/retry"
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
		time.Second,
		time.Millisecond,
		func() error {
			calls++

			if calls < 3 {
				return errors.New("down")
			}

			return nil
		},
	)
	assert.Nil(t, e)
	assert.Integer(t, 3, r.Attempts)
	assert.True(t, r.Elapsed >= 2*time.Millisecond)
}

func TestUntilGivesUpAtTheDeadline(t *testing.T) {
	fail := errors.New("down")
	r, e := retry.Until(
		10*time.Millisecond,
		time.Millisecond,
		func() error {
			return fail
		},
	)
	assert.True(t, r.Attempts > 1)
	assert.True(t, r.Elapsed < 20*time.Millisecond)
	assert.StringContains(t, "after 10ms", e.Error())
	assert.True(t, errors.Is(e, fail))
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
