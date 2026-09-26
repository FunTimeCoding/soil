package unit

func policyOutput() string {
	return `foxtrot:
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
}
