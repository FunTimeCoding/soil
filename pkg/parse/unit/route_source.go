package unit

func routeSource() string {
	return "package test\n\nfunc Run() {\n\tm.HandleFunc(\"GET /palette\", a)\n\tm.HandleFunc(\"GET /favicon.ico\", b)\n}\n"
}
