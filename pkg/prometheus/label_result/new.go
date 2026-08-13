package label_result

func New(
	v []string,
	w []string,
) *Result {
	return &Result{Values: v, Warnings: w}
}
