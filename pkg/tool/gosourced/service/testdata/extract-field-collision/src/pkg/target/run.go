package target

func Run() string {
	v := Store{name: "alfa", Name: "bravo"}

	return v.name + v.Name
}
