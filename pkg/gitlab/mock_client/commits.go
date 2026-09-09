package mock_client

func (c *Client) Commits() []*RecordedCommit {
	return c.commits
}
