package mock_notifier

func (n *Notifier) Reset() {
	n.mutex.Lock()
	defer n.mutex.Unlock()
	n.Notified = 0
}
