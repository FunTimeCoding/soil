package service_tester

import "github.com/funtimecoding/soil/pkg/nextcloud/face"

func (o *Tester) usage() face.UsageClient {
	if o.Usage == nil {
		return nil
	}

	return o.Usage
}
