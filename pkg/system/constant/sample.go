package constant

const (
	FixtureEchoText  = "test"
	FixtureDirectory = "test"

	FixtureAptPolicy = `foxtrot:
  Installed: 1.2.3
  Candidate: 1.2.5
  Version table:
     1.2.5 500
        500 https://apt.example.org stable/main amd64 Packages
 *** 1.2.3 500
        500 https://apt.example.org stable/main amd64 Packages
        100 /var/lib/dpkg/status
     1.2.0 500
        500 https://apt.example.org stable/main amd64 Packages
avahi-daemon:
  Installed: 0.8-16
  Candidate: 0.8-16
  Version table:
 *** 0.8-16 500
        500 mirror+file:/etc/apt/mirrors/debian.list trixie/main amd64 Packages
        100 /var/lib/dpkg/status
`

	FixtureDpkgSearch = `avahi-daemon: /usr/lib/systemd/system/avahi-daemon.service
foxtrot: /lib/systemd/system/foxtrot.service
dpkg-query: no path found matching pattern /etc/systemd/system/custom.service
`

	FixtureSystemctlUnits = `avahi-daemon.service    loaded    active   running Avahi mDNS/DNS-SD Stack
foxtrot.service         loaded    failed   failed  foxtrot stub description
ssh.service             loaded    active   running OpenBSD Secure Shell server
fstrim.service          loaded    inactive dead    Discard unused filesystem blocks
NetworkManager.service  not-found inactive dead    NetworkManager.service
`
)
