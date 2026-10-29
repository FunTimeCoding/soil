package mattermost

import "fmt"

func emptyUpload(name string) error {
	return fmt.Errorf("upload %s returned no file info", name)
}
