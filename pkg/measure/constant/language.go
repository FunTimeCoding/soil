package constant

import (
	soil "github.com/funtimecoding/soil/pkg/constant"
	go_modConstant "github.com/funtimecoding/soil/pkg/go_mod/constant"
	"github.com/funtimecoding/soil/pkg/measure/language"
	stringsConstant "github.com/funtimecoding/soil/pkg/strings/constant"
)

var (
	GoTest = language.New("Go Test").
		WithSuffix(soil.TestSuffix).
		WithLineComment(stringsConstant.DoubleSlash).
		WithBlockComment(BlockOpenC, BlockCloseC).
		WithQuote(QuoteDouble).
		WithRawQuote(QuoteBacktick)
	Go = language.New("Go").
		WithExtension(soil.GoExtension).
		WithLineComment(stringsConstant.DoubleSlash).
		WithBlockComment(BlockOpenC, BlockCloseC).
		WithQuote(QuoteDouble).
		WithRawQuote(QuoteBacktick)
	GoModule = language.New("Go Module").
			WithFilename(go_modConstant.ModFile, go_modConstant.SumFile).
			WithLineComment(stringsConstant.DoubleSlash)
	C = language.New("C").
		WithExtension(".c", ".h").
		WithLineComment(stringsConstant.DoubleSlash).
		WithBlockComment(BlockOpenC, BlockCloseC).
		WithQuote(QuoteDouble)
	CPlusPlus = language.New("C++").
			WithExtension(".cc", ".cpp", ".cxx", ".hpp", ".hh").
			WithLineComment(stringsConstant.DoubleSlash).
			WithBlockComment(BlockOpenC, BlockCloseC).
			WithQuote(QuoteDouble)
	Java = language.New("Java").
		WithExtension(".java").
		WithLineComment(stringsConstant.DoubleSlash).
		WithBlockComment(BlockOpenC, BlockCloseC).
		WithQuote(QuoteDouble)
	Kotlin = language.New("Kotlin").
		WithExtension(".kt", ".kts").
		WithLineComment(stringsConstant.DoubleSlash).
		WithBlockComment(BlockOpenC, BlockCloseC).
		WithQuote(QuoteDouble)
	Rust = language.New("Rust").
		WithExtension(".rs").
		WithLineComment(stringsConstant.DoubleSlash).
		WithBlockComment(BlockOpenC, BlockCloseC).
		WithQuote(QuoteDouble).
		WithNested()
	JavaScript = language.New("JavaScript").
			WithExtension(".js", ".mjs", ".cjs", ".jsx").
			WithShebang("node").
			WithLineComment(stringsConstant.DoubleSlash).
			WithBlockComment(BlockOpenC, BlockCloseC).
			WithQuote(QuoteDouble, QuoteSingle).
			WithRawQuote(QuoteBacktick)
	TypeScript = language.New("TypeScript").
			WithExtension(".ts", ".tsx").
			WithLineComment(stringsConstant.DoubleSlash).
			WithBlockComment(BlockOpenC, BlockCloseC).
			WithQuote(QuoteDouble, QuoteSingle).
			WithRawQuote(QuoteBacktick)
	Vue = language.New("Vue").
		WithExtension(".vue").
		WithLineComment(stringsConstant.DoubleSlash).
		WithBlockComment(BlockOpenC, BlockCloseC).
		WithBlockComment(BlockOpenHypertext, BlockCloseHypertext)
	PHP = language.New("PHP").
		WithExtension(".php", ".phtml").
		WithShebang("php").
		WithLineComment(stringsConstant.DoubleSlash, stringsConstant.Hash).
		WithBlockComment(BlockOpenC, BlockCloseC).
		WithQuote(QuoteDouble, QuoteSingle)
	Swift = language.New("Swift").
		WithExtension(".swift").
		WithLineComment(stringsConstant.DoubleSlash).
		WithBlockComment(BlockOpenC, BlockCloseC).
		WithQuote(QuoteDouble).
		WithNested()
	Scala = language.New("Scala").
		WithExtension(".scala", ".sc").
		WithLineComment(stringsConstant.DoubleSlash).
		WithBlockComment(BlockOpenC, BlockCloseC).
		WithQuote(QuoteDouble).
		WithNested()
	Groovy = language.New("Groovy").
		WithExtension(".groovy", ".gradle").
		WithFilename("Jenkinsfile").
		WithLineComment(stringsConstant.DoubleSlash).
		WithBlockComment(BlockOpenC, BlockCloseC).
		WithQuote(QuoteDouble, QuoteSingle)
	CSharp = language.New("C#").
		WithExtension(".cs").
		WithLineComment(stringsConstant.DoubleSlash).
		WithBlockComment(BlockOpenC, BlockCloseC).
		WithQuote(QuoteDouble)
	Protobuf = language.New("Protocol Buffers").
			WithExtension(".proto").
			WithLineComment(stringsConstant.DoubleSlash).
			WithBlockComment(BlockOpenC, BlockCloseC)
	Terraform = language.New("Terraform").
			WithExtension(".tf", ".tfvars", ".hcl").
			WithLineComment(stringsConstant.Hash, stringsConstant.DoubleSlash).
			WithBlockComment(BlockOpenC, BlockCloseC).
			WithQuote(QuoteDouble)
	CSS = language.New("CSS").
		WithExtension(".css").
		WithBlockComment(BlockOpenC, BlockCloseC)
	Less = language.New("Less").
		WithExtension(".less").
		WithLineComment(stringsConstant.DoubleSlash).
		WithBlockComment(BlockOpenC, BlockCloseC)
	Sass = language.New("Sass").
		WithExtension(".scss", ".sass").
		WithLineComment(stringsConstant.DoubleSlash).
		WithBlockComment(BlockOpenC, BlockCloseC)

	Shell = language.New("Shell").
		WithExtension(".sh", ".bash", ".zsh", ".ksh").
		WithFilename(
			".bashrc",
			".bash_profile",
			".zshrc",
			".profile",
			"PKGBUILD",
		).
		WithShebang("sh", "bash", "zsh", "ksh", "dash", "ash").
		WithLineComment(stringsConstant.Hash)
	Fish = language.New("Fish").
		WithExtension(".fish").
		WithShebang("fish").
		WithLineComment(stringsConstant.Hash)
	PowerShell = language.New("PowerShell").
			WithExtension(".ps1", ".psm1", ".psd1").
			WithShebang("pwsh").
			WithLineComment(stringsConstant.Hash).
			WithBlockComment("<#", "#>")
	Python = language.New("Python").
		WithExtension(".py", ".pyi", ".pyw").
		WithShebang("python").
		WithLineComment(stringsConstant.Hash).
		WithBlockComment(`"""`, `"""`).
		WithBlockComment("'''", "'''")
	Ruby = language.New("Ruby").
		WithExtension(".rb", ".rake", ".gemspec").
		WithFilename(
			"Gemfile",
			"Rakefile",
			"Berksfile",
			"Vagrantfile",
			"Guardfile",
			"Podfile",
			"Capfile",
		).
		WithShebang("ruby").
		WithLineComment(stringsConstant.Hash).
		WithBlockComment("=begin", "=end")
	Perl = language.New("Perl").
		WithExtension(".pl", ".pm", ".t").
		WithShebang("perl").
		WithLineComment(stringsConstant.Hash).
		WithBlockComment("=pod", "=cut")
	Makefile = language.New("Makefile").
			WithExtension(".mk").
			WithFilename("Makefile", "makefile", "GNUmakefile").
			WithLineComment(stringsConstant.Hash)
	Dockerfile = language.New("Dockerfile").
			WithFilename(soil.DockerFile, soil.ContainerFile).
			WithExtension(".dockerfile").
			WithLineComment(stringsConstant.Hash)
	Markup = language.New("YAML").
		WithExtension(soil.MarkupExtension, soil.ShortMarkupExtension).
		WithLineComment(stringsConstant.Hash)
	Salt = language.New("Salt").
		WithExtension(".sls").
		WithLineComment(stringsConstant.Hash).
		WithBlockComment("{#", "#}")
	Jinja = language.New("Jinja").
		WithExtension(".j2", ".jinja", ".jinja2").
		WithBlockComment("{#", "#}")
	TOML = language.New("TOML").
		WithExtension(".toml").
		WithLineComment(stringsConstant.Hash)
	INI = language.New("INI").
		WithExtension(".ini", ".cfg", ".conf", ".cnf", ".properties").
		WithLineComment(stringsConstant.Hash, stringsConstant.Semicolon)
	Systemd = language.New("Systemd").
		WithExtension(".service", ".timer", ".socket", ".mount", ".target").
		WithLineComment(stringsConstant.Hash, stringsConstant.Semicolon)
	Cron = language.New("Cron").
		WithExtension(".cron").
		WithFilename("crontab").
		WithLineComment(stringsConstant.Hash)
	GitConfiguration = language.New("Git Configuration").
				WithFilename(
			".gitignore",
			".gitattributes",
			".gitmodules",
			".dockerignore",
			".helmignore",
			"chefignore",
		).
		WithLineComment(stringsConstant.Hash)
	R = language.New("R").
		WithExtension(".r", ".R").
		WithLineComment(stringsConstant.Hash)
	Elixir = language.New("Elixir").
		WithExtension(".ex", ".exs").
		WithLineComment(stringsConstant.Hash)

	Lua = language.New("Lua").
		WithExtension(".lua").
		WithShebang("lua").
		WithLineComment("--").
		WithBlockComment("--[[", "]]")
	SQL = language.New("SQL").
		WithExtension(".sql").
		WithLineComment("--").
		WithBlockComment(BlockOpenC, BlockCloseC).
		WithQuote(QuoteSingle)
	Haskell = language.New("Haskell").
		WithExtension(".hs", ".lhs").
		WithLineComment("--").
		WithBlockComment("{-", "-}").
		WithQuote(QuoteDouble).
		WithNested()
	Elm = language.New("Elm").
		WithExtension(".elm").
		WithLineComment("--").
		WithBlockComment("{-", "-}").
		WithQuote(QuoteDouble).
		WithNested()

	Lisp = language.New("Lisp").
		WithExtension(
			".lisp",
			".cl",
			".el",
			".clj",
			".cljs",
			".edn",
			".scm",
			".rkt",
		).
		WithLineComment(stringsConstant.Semicolon).
		WithBlockComment("#|", "|#")
	Assembly = language.New("Assembly").
			WithExtension(".asm", ".s", ".S").
			WithLineComment(stringsConstant.Semicolon, stringsConstant.Hash).
			WithBlockComment(BlockOpenC, BlockCloseC)

	Erlang = language.New("Erlang").
		WithExtension(".erl", ".hrl").
		WithLineComment("%")
	LaTeX = language.New("LaTeX").
		WithExtension(".tex", ".sty", ".cls", ".bib").
		WithLineComment("%")
	Prolog = language.New("Prolog").
		WithExtension(".pro", ".P").
		WithLineComment("%").
		WithBlockComment(BlockOpenC, BlockCloseC)

	Hypertext = language.New("HTML").
			WithExtension(soil.HypertextExtension, ".htm", ".xhtml").
			WithBlockComment(BlockOpenHypertext, BlockCloseHypertext)
	ExtensibleMarkup = language.New("XML").
				WithExtension(
			".xml",
			".xsd",
			".xsl",
			".xslt",
			".svg",
			".plist",
			".csproj",
			".avsc",
		).
		WithBlockComment(BlockOpenHypertext, BlockCloseHypertext)
	Markdown = language.New("Markdown").
			WithExtension(soil.MarkdownExtension, ".markdown").
			WithBlockComment(BlockOpenHypertext, BlockCloseHypertext)
	ERB = language.New("ERB").
		WithExtension(".erb").
		WithBlockComment("<%#", "%>").
		WithBlockComment(BlockOpenHypertext, BlockCloseHypertext)
	GoTemplate = language.New("Go Template").
			WithExtension(".tmpl", ".tpl", ".gotmpl").
			WithBlockComment("{{/*", "*/}}").
			WithBlockComment(BlockOpenHypertext, BlockCloseHypertext)

	Notation = language.New("JSON").
			WithExtension(".json", ".jsonl", ".geojson", ".webmanifest")
	CSV = language.New("CSV").
		WithExtension(".csv", ".tsv")
	PlainText = language.New("Plain Text").
			WithExtension(".txt", ".text", ".list").
			WithFilename("LICENSE", "COPYING", "AUTHORS", "CHANGELOG", "NOTICE")
	OpenSSH = language.New("OpenSSH Key").
		WithExtension(".pub")

	Languages = []*language.Language{
		GoTest,
		Go,
		GoModule,
		C,
		CPlusPlus,
		Java,
		Kotlin,
		Rust,
		JavaScript,
		TypeScript,
		Vue,
		PHP,
		Swift,
		Scala,
		Groovy,
		CSharp,
		Protobuf,
		Terraform,
		CSS,
		Less,
		Sass,
		Shell,
		Fish,
		PowerShell,
		Python,
		Ruby,
		Perl,
		Makefile,
		Dockerfile,
		Markup,
		Salt,
		Jinja,
		TOML,
		INI,
		Systemd,
		Cron,
		GitConfiguration,
		R,
		Elixir,
		Lua,
		SQL,
		Haskell,
		Elm,
		Lisp,
		Assembly,
		Erlang,
		LaTeX,
		Prolog,
		Hypertext,
		ExtensibleMarkup,
		Markdown,
		ERB,
		GoTemplate,
		Notation,
		CSV,
		PlainText,
		OpenSSH,
	}
)
