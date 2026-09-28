package mock_client

func (c *Client) Play(
	sessionIdentifier string,
	itemIdentifiers []string,
	playCommand string,
	startPositionTicks int64,
) error {
	return nil
}
