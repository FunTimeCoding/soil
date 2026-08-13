package pod

func New(
	namespace string,
	name string,
	phase string,
) *Pod {
	return &Pod{Namespace: namespace, Name: name, Phase: phase}
}
