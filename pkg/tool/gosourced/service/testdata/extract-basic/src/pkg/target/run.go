package target

func Run() string {
	v := Store{name: "alfa"}

	return v.save()
}
