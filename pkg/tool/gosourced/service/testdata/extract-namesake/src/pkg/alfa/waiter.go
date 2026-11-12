package alfa

type waiter struct {
	ready   chan struct{}
	granted bool
}
