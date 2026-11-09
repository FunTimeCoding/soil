package part

type Foxtrot struct{}

func (f *Foxtrot) Size() int {
	return 6
}

func Total() int {
	d := Delta{}

	return (&Alfa{}).Size() + Bravo{}.Size() + d.Size() + Echo(4).Size()
}
