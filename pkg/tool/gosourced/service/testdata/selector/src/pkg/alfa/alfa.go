package alfa

type Alfa struct{}

func (a *Alfa) Run(
	x int,
	y int,
) int {
	return x + y
}
