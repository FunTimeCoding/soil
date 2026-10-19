package outage

import (
	"fmt"
	"time"
)

func (o *OutageError) Error() string {
	return fmt.Sprintf(
		"%s down for %s: %v",
		o.Peer,
		o.Duration.Round(time.Second),
		o.cause,
	)
}
