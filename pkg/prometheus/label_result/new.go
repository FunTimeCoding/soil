package label_result

func New(
	v []string,
	w []string,
	n []string,
) *Result {
	return &Result{Values: v, Warnings: w, Notices: n}
}
