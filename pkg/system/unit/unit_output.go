package unit

func unitOutput() string {
	return `avahi-daemon.service    loaded    active   running Avahi mDNS/DNS-SD Stack
foxtrot.service         loaded    failed   failed  foxtrot stub description
ssh.service             loaded    active   running OpenBSD Secure Shell server
fstrim.service          loaded    inactive dead    Discard unused filesystem blocks
NetworkManager.service  not-found inactive dead    NetworkManager.service
`
}
