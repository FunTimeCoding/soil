package request

type Call struct {
	Declaration  string  `json:"functionDeclaration"`
	Target       *Target `json:"target"`
	Arguments    []any   `json:"arguments,omitempty"`
	AwaitPromise bool    `json:"awaitPromise"`
}
