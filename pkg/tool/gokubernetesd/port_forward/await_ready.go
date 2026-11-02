package port_forward

func AwaitReady(
	ready <-chan struct{},
	failed <-chan error,
) error {
	select {
	case <-ready:
		return nil
	case e := <-failed:
		return e
	}
}
