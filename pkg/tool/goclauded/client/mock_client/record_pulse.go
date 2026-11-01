package mock_client

import "github.com/funtimecoding/soil/pkg/tool/goclauded/types/pulse_call"

func (c *Client) recordPulse(
	sessionIdentifier string,
	body string,
) (bool, error) {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	if c.Failure != nil {
		return false, c.Failure
	}

	c.Pulses = append(c.Pulses, pulse_call.New(sessionIdentifier, body))

	return false, nil
}
