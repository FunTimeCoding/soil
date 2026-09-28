package mock_client

func (c *Client) PlaybackCommand(
	sessionIdentifier string,
	command string,
	seekPositionTicks int64,
) error {
	return nil
}
