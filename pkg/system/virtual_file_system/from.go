package virtual_file_system

func From(directory string) *System {
	return FromSkipping(directory, func(string) bool { return false })
}
