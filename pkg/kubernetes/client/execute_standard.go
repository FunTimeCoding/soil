package client

import "os"

func (c *Client) ExecuteStandard(
	namespace string,
	pod string,
	container string,
	command ...string,
) error {
	return c.Execute(
		os.Stdout,
		os.Stderr,
		namespace,
		pod,
		container,
		command...,
	)
}
