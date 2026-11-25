package watched_request

import "fmt"

func (k *Key) String() string {
	return fmt.Sprintf("%s/%d!%d", k.Instance, k.Project, k.Request)
}
