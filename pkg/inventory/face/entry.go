package face

type Entry[T any] interface {
	*T
	InstanceName() string
	InstanceHost() string
}
