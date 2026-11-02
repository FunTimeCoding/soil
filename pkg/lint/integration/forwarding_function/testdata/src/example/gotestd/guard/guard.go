package guard

type Guard struct{}

func (g *Guard) Wait(timeout int) error {
	if timeout < 0 {
		return nil
	}

	return nil
}
