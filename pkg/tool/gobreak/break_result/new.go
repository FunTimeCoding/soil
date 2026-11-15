package break_result

func New(
	name string,
	expect string,
	status string,
) *Result {
	return &Result{Name: name, Expect: expect, Status: status}
}
