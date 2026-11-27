package constant

const (
	ChromiumHostEnvironment = "CHROMIUM_HOST"
	ChromiumPortEnvironment = "CHROMIUM_PORT"
	ChromiumDefaultPort     = 9222

	ChromiumBodySelector = "body"

	ChromiumPageTabType   = "page"
	ChromiumIframeTabType = "iframe"

	ChromiumNotationPath = "/json"
	ChromiumVersionPath  = "/json/version"

	ChromiumTabNotFound = "tab not found"
	ChromiumBrowserTab  = "browser tab"
	ChromiumAsleep      = "did not answer within %s - a sleeping tab answers once activated"
)
