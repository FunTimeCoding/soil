package mock

type Mock struct{}

func (m *Mock) Size() int {
	return 2
}

func (m *Mock) Unit() string {
	return "cm"
}
