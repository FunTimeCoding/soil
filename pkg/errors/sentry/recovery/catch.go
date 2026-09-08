package recovery

func Catch(f func()) (v any) {
	defer func() { v = recover() }()
	f()

	return v
}
