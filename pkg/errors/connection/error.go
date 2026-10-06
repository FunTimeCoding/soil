package connection

import "fmt"

func (f *Failure) Error() string {
	if f.Host == "" {
		return f.Reason
	}

	return fmt.Sprintf("%s: %s", f.Host, f.Reason)
}
