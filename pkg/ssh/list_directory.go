package ssh

import "io/fs"

func (c *Client) ListDirectory(path string) ([]fs.FileInfo, error) {
	s, e := c.sftpClient()

	if e != nil {
		return nil, e
	}

	result, e := s.ReadDir(path)

	if e != nil {
		return nil, c.refused("list directory", path, e)
	}

	return result, nil
}
