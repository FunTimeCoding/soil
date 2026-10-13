package requester

func (r *Requester) Close() {
	r.client.CloseIdleConnections()
}
