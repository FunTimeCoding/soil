package turn

type waiter struct {
	ready   chan struct{}
	granted bool
}
