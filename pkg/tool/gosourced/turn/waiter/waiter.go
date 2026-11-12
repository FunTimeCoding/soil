package waiter

type Waiter struct {
	Ready   chan struct{}
	Granted bool
}
