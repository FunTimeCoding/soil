package recovering_discovery

import (
	"errors"
	"k8s.io/apimachinery/pkg/apis/meta/v1"
)

func (d *Discovery) ServerPreferredResources() ([]*v1.APIResourceList, error) {
	d.mutex.Lock()
	defer d.mutex.Unlock()

	if d.failures > 0 {
		d.failures--

		return nil, errors.New("discovery unavailable")
	}

	return d.resources, nil
}
