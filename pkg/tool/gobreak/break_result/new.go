package break_result

func New(
	name string,
	expect string,
	kind string,
) *Result {
	return &Result{Name: name, Expect: expect, Kind: kind}
}
