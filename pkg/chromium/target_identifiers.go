package chromium

import "sort"

func (c *Client) TargetIdentifiers() []string {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	result := make([]string, 0, len(c.targets))

	for identifier := range c.targets {
		result = append(result, identifier)
	}

	sort.Strings(result)

	return result
}
