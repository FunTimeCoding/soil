package service_tester

import "errors"

func Fault() error {
	return errors.New("mattermost answered garbage")
}
