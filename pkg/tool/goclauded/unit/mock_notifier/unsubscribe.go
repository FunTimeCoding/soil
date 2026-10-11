package mock_notifier

func (n *Notifier) Unsubscribe(c chan struct{}) {
	n.mutex.Lock()
	defer n.mutex.Unlock()
	delete(n.subscribers, c)
	close(c)
}
