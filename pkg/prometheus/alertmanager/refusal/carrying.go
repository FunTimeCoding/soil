package refusal

type carrying interface {
	error
	GetPayload() string
}
