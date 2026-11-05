package mock_client

func (c *Client) DeleteJournalEntry(_ int32) error {
	if c.Failure != nil {
		return c.Failure
	}

	return nil
}
