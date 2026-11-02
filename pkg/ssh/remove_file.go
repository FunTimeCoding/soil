package ssh

func (c *Client) RemoveFile(path string) error {
	s, e := c.sftpClient()

	if e != nil {
		return e
	}

	if f := s.Remove(path); f != nil {
		return c.refused("remove", path, f)
	}

	return nil
}
