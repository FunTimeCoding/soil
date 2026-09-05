package mock_notifier

func (n *Notifier) Subscribe() chan struct{} {
	n.mutex.Lock()
	defer n.mutex.Unlock()
	c := make(chan struct{}, 1)
	n.subscribers[c] = struct{}{}

	return c
}
