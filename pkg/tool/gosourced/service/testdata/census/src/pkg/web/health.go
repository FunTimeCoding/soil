package web

import "example/pkg/client"

func Healthy(c *client.Client) bool {
	if c.Ready() {
		return true
	}

	return false
}

func Gate(c *client.Client) string {
	if c.Ready() {
		return "open"
	}

	return "closed"
}

func Probe(c *client.Client, retries int) bool {
	if c.Ready() && retries < 3 {
		return true
	}

	return false
}
