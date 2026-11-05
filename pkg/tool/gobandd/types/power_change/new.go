package power_change

func New(
	power string,
	requested string,
) *Change {
	return &Change{Power: power, Requested: requested}
}
