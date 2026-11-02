package ambiguous_pods

func New(
	name string,
	pods []string,
) *AmbiguousPods {
	return &AmbiguousPods{Name: name, Pods: pods}
}
