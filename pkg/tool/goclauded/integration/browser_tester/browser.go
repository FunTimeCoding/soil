package browser_tester

import (
	"context"
	"testing"
)

type Browser struct {
	T       *testing.T
	Context context.Context
	cancel  context.CancelFunc
}
