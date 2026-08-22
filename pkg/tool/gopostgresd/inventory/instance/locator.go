package instance

import "fmt"

func (i *Instance) Locator() string {
	return fmt.Sprintf(
		"postgres://%s:%s@%s:%d/%s",
		i.User,
		i.Password,
		i.Host,
		i.Port,
		i.Database,
	)
}
