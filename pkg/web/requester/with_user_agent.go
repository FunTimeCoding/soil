package requester

func (r *Requester) WithUserAgent(agent string) *Requester {
	r.userAgent = agent

	return r
}
