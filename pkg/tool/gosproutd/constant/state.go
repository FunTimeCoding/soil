package constant

type State string

const (
	StateOpen       State = "open"
	StateAnswered   State = "answered"
	StateDeclined   State = "declined"
	StateIrrelevant State = "irrelevant"
	StatePostponed  State = "postponed"
)
