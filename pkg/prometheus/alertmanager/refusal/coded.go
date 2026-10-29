package refusal

type coded interface {
	error
	Code() int
}
