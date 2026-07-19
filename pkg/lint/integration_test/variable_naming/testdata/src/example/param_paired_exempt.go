package example

type IssueValue struct{}

func ParamPairedExempt(
	before *IssueValue,
	after *IssueValue,
) {
	_ = before
	_ = after
}
