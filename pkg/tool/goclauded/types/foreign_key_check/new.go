package foreign_key_check

func New(
	child string,
	column string,
	parent string,
) *Check {
	return &Check{Child: child, Column: column, Parent: parent}
}
