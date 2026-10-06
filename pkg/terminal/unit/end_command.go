package unit

func (c *commandEnds) EndCommand(outcome string) {
	c.outcomes = append(c.outcomes, outcome)
}
