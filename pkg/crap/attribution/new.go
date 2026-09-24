package attribution

func New() *Matrix {
	return &Matrix{Covers: map[string][]string{}}
}
