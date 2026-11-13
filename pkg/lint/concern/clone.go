package concern

func (c *Concern) Clone() *Concern {
	result := *c

	return &result
}
