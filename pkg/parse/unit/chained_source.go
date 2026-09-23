package unit

func chainedSource() string {
	return "package test\n\nfunc Run() {\n\tlayout.New(i).WithTheme(theme.Straw).WithCommandPalette(\"/palette\")\n}\n"
}
