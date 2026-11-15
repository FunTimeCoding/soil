package assistant_server

func signal(c chan struct{}) {
	select {
	case c <- struct{}{}:
	default:
	}
}
