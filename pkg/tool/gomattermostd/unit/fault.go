package unit

import "errors"

func fault() error {
	return errors.New("mattermost answered garbage")
}
