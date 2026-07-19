package source

func Helper() string {
	return "help"
}

func Mover() string {
	return Helper()
}

func Stayer() string {
	return Mover() + " stays"
}
