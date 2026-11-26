package request

type Evaluate struct {
	Expression   string  `json:"expression"`
	Target       *Target `json:"target"`
	AwaitPromise bool    `json:"awaitPromise"`
}
