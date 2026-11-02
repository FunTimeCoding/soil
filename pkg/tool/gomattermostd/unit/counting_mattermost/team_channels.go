package counting_mattermost

func (c *Client) TeamChannels() int {
	return int(c.teamChannel.Load())
}
