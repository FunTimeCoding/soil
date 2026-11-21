package kernel_offer

func New(
	version string,
	priority int,
	origins []string,
) *Offer {
	return &Offer{Version: version, Priority: priority, Origins: origins}
}
