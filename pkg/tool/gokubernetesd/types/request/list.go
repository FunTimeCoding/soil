package request

type List struct {
	ResourceType  string
	Namespace     string
	AllNamespaces bool
	LabelSelector string
	FieldSelector string
}
