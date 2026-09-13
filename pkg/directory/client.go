package directory

type Client struct {
	host        string
	port        int
	base        string
	bind        string
	password    string
	userFilter  string
	groupFilter string
	authority   string
	insecure    bool
	untrusted   bool
}
