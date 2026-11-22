package callback_report

type Outcome struct {
	ReturnCode  int    `json:"rc"`
	Stdout      string `json:"stdout"`
	Stderr      string `json:"stderr"`
	Failed      bool   `json:"failed"`
	Unreachable bool   `json:"unreachable"`
	Message     string `json:"msg"`
}
