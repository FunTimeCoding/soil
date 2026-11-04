package trace_transport

func (t *Transport) CloseIdleConnections() {
	t.transport.CloseIdleConnections()
}
