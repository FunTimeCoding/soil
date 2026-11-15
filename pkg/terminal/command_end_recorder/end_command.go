package command_end_recorder

func (c *Recorder) EndCommand(outcome string) {
	c.Outcomes = append(c.Outcomes, outcome)
}
