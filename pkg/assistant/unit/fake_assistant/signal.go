package fake_assistant

func signal(c chan struct{}) {
	select {
	case c <- struct{}{}:
	default:
	}
}
