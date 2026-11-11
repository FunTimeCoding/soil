package source

func stayer() string {
	return "stay"
}

func Stayer() string {
	return "Stay"
}

func Mover() string {
	return stayer() + " moved"
}

func Keeper() string {
	return Stayer()
}
