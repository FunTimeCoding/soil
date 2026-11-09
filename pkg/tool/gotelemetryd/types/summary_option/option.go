package summary_option

import "time"

type Option struct {
	Since   string
	Until   string
	Span    time.Duration
	GroupBy string
	Tool    string
	Actor   string
}
