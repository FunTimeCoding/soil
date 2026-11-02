package argument

type Timeline struct {
	Since  string
	Before string
	Kinds  []string
	Limit  int
	Offset int
	Full   bool
}
