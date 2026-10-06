package matrix

type Client struct {
	Path     string
	Repo     string
	Must     bool
	Basic    bool
	Entity   bool
	Constant bool
	Example  bool
}
