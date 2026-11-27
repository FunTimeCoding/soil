package constant

import "fmt"

const (
	ChromiumFixtureQuietPage = `<!doctype html><title>quiet</title>
<body><p id="mark">quiet</p></body>`

	ChromiumFixtureUnloadPage = `<!doctype html><title>unload</title>
<body><p id="mark">unload</p><script>
window.addEventListener('beforeunload', function (e) {
  e.preventDefault();
  e.returnValue = '';
});
</script></body>`

	ChromiumFixtureHeavyPagePrefix = `<!doctype html><title>heavy</title>
<body><p id="mark">heavy</p>`

	ChromiumFixtureHeavyPageSuffix = `</body>`

	ChromiumFixtureHeavyRowFormat = "<div class=\"row\">row %d padding text</div>"
	busyPageFormat                = `<!doctype html><title>busy</title>
<body><p id="mark">busy</p><div id="sink"></div><script>
setInterval(function () {
  var d = document.createElement('span');
  d.textContent = String(Date.now());
  var sink = document.getElementById('sink');
  sink.appendChild(d);

  if (sink.childNodes.length > 40) {
    sink.removeChild(sink.firstChild);
  }

  fetch('%s').catch(function () {});
}, 50);
</script></body>`

	stalledPageFormat = `<!doctype html><title>stalled</title>
<body><p id="mark">stalled</p>
<img src="%s" alt="never arrives"></body>`

	popupPageFormat = `<!doctype html><title>popup</title>
<body><p id="mark">popup</p>
<script>window.open('%s', '_blank');</script></body>`
)

var (
	ChromiumFixtureBusyPage = fmt.Sprintf(
		busyPageFormat,
		ChromiumFixturePingRoute,
	)
	ChromiumFixtureStalledPage = fmt.Sprintf(
		stalledPageFormat,
		ChromiumFixtureHangRoute,
	)
	ChromiumFixturePopupPage = fmt.Sprintf(
		popupPageFormat,
		ChromiumFixtureQuietRoute,
	)
)
