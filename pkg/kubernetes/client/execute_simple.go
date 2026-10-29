package client

import "strings"

func (c *Client) ExecuteSimple(
	namespace string,
	pod string,
	command ...string,
) (string, string, error) {
	stdout := &strings.Builder{}
	stderr := &strings.Builder{}
	e := c.Execute(stdout, stderr, namespace, pod, "", command...)

	return stdout.String(), stderr.String(), e
}
