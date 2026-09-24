package configuration

import "slices"

func (c *Configuration) HookNames() []string {
	result := make([]string, 0, len(c.Hooks))

	for name := range c.Hooks {
		result = append(result, name)
	}

	slices.Sort(result)

	return result
}
