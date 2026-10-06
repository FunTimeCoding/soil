package constant

import "time"

const (
	Retries           = 2
	Backoff           = time.Second
	MaximumRetryAfter = 30 * time.Second
	RetryAfter        = "Retry-After"
	MaximumRefusal    = 64 << 10
)

var NestedFields = []string{"message", "title", "detail"}
var ReasonFields = []string{
	"reason",
	"message",
	"detail",
	"error_description",
	"error",
	"errorMessages",
	"errorMessage",
	"errors",
	"title",
}
