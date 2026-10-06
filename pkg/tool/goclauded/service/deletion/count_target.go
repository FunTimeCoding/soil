package deletion

type CountTarget struct {
	Read   func(string) (int64, error)
	Target *int64
}
