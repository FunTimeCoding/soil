package constant

type AnswerKind string

const (
	AnswerKindChoice     AnswerKind = "choice"
	AnswerKindOther      AnswerKind = "other"
	AnswerKindConstraint AnswerKind = "constraint"
)
