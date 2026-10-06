package apply_result

func New(
	kind string,
	name string,
	namespace string,
) *Result {
	return &Result{Kind: kind, Name: name, Namespace: namespace}
}
