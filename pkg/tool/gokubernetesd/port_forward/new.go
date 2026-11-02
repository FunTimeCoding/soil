package port_forward

func New(
	identifier string,
	resource string,
	namespace string,
	localPort int,
	targetPort int,
) *Forward {
	return &Forward{
		Identifier: identifier,
		Resource:   resource,
		Namespace:  namespace,
		LocalPort:  localPort,
		TargetPort: targetPort,
		stop:       make(chan struct{}),
	}
}
