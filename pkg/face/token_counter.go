package face

type TokenCounter interface {
	Count(text string) int
	Allowance() int
}
