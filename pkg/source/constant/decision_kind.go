package constant

const (
	DecisionKeep DecisionKind = iota
	DecisionRename
	DecisionRefused
	DecisionUnderivable
)

type DecisionKind int
