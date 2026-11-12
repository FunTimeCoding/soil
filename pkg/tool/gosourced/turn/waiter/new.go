package waiter

func New(ready chan struct{}) *Waiter {
	return &Waiter{Ready: ready}
}
