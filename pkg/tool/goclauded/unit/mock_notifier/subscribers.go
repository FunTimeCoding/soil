package mock_notifier

func (n *Notifier) Subscribers() int {
	n.mutex.Lock()
	defer n.mutex.Unlock()

	return len(n.subscribers)
}
