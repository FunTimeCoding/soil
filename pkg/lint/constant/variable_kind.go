package constant

const (
	VariableKindLocal VariableKind = iota
	VariableKindParameter
	VariableKindReceiver
)

type VariableKind int
