package mock_client

func (c *Client) Replies(thread string) []string {
	var result []string

	for _, v := range c.threads[thread] {
		result = append(result, v.Body)
	}

	return result
}
