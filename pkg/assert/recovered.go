package assert

func recovered(f func()) (result any, panicked bool) {
	panicked = true

	defer func() {
		if panicked {
			result = recover()
		}
	}()
	f()
	panicked = false

	return result, panicked
}
