package client

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/kubernetes/constant"
	"k8s.io/apimachinery/pkg/apis/meta/v1"
)

func (c *Client) DeleteCertificateRequest(
	namespace string,
	name string,
) {
	errors.PanicOnError(
		c.dynamic.Resource(constant.CertificateRequestGVR).Namespace(namespace).Delete(
			c.context,
			name,
			v1.DeleteOptions{},
		),
	)
}
