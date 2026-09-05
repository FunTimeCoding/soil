package mock_notifier

func (n *Notifier) Notify() {
	n.mutex.Lock()
	defer n.mutex.Unlock()
	n.Notified++

	for c := range n.subscribers {
		select {
		case c <- struct{}{}:
		default:
		}
	}
}
