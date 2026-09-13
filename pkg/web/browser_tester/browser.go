package browser_tester

import (
	"context"
	"testing"
	"time"
)

type Browser struct {
	T       *testing.T
	Context context.Context
	Timeout time.Duration
	cancel  context.CancelFunc
}
