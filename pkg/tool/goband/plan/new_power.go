package plan

func NewPower(
	instance string,
	power string,
) *Power {
	return &Power{Instance: instance, Power: power}
}
