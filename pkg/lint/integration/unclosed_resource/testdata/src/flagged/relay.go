package example

type Settings struct {
	Make func() *Client
}

func newSettings(c *Client) *Settings {
	result := &Settings{}
	result.Make = func() *Client { return c }

	return result
}

type Dropper struct{}

func newDropper(s *Settings) *Dropper {
	s.Make().Work()

	return &Dropper{}
}

func RelayedAndDropped() {
	_ = newDropper(newSettings(newClient()))
}
