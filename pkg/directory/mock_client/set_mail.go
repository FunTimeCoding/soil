package mock_client

func (c *Client) SetMail(
	unique string,
	mail string,
) {
	for _, n := range c.entries {
		if n.Unique == unique {
			n.Mail = mail
		}
	}
}
