package integration

import (
	"context"
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/system/run"
	"testing"
	"time"
)

func TestRunOutlivingItsContextIsStopped(t *testing.T) {
	windowsSkip(t)
	x, cancel := context.WithTimeout(context.Background(), 50*time.Millisecond)
	defer cancel()
	began := time.Now()
	r := run.New().NoPanic().WithContext(x)
	r.Start("sleep", "5")
	assert.True(t, r.Error != nil)
	assert.True(t, time.Since(began) < 4*time.Second)
}

func TestRunWithinItsContextFinishes(t *testing.T) {
	windowsSkip(t)
	x, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	r := run.New().NoPanic().WithContext(x)
	assert.String(t, "ready", r.Start("echo", "ready"))
	assert.True(t, r.Error == nil)
}
