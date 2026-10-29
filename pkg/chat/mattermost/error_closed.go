package mattermost

import "errors"

func errorClosed() error {
	return errors.New("mattermost client closed")
}
