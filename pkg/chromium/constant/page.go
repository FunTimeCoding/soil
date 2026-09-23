package constant

const (
	FixtureQuietPage = `<!doctype html><title>quiet</title>
<body><p id="mark">quiet</p></body>`

	FixtureBusyPage = `<!doctype html><title>busy</title>
<body><p id="mark">busy</p><div id="sink"></div><script>
setInterval(function () {
  var d = document.createElement('span');
  d.textContent = String(Date.now());
  var sink = document.getElementById('sink');
  sink.appendChild(d);

  if (sink.childNodes.length > 40) {
    sink.removeChild(sink.firstChild);
  }

  fetch('` + FixturePingRoute + `').catch(function () {});
}, 50);
</script></body>`

	FixtureStalledPage = `<!doctype html><title>stalled</title>
<body><p id="mark">stalled</p>
<img src="` + FixtureHangRoute + `" alt="never arrives"></body>`

	FixtureUnloadPage = `<!doctype html><title>unload</title>
<body><p id="mark">unload</p><script>
window.addEventListener('beforeunload', function (e) {
  e.preventDefault();
  e.returnValue = '';
});
</script></body>`

	FixturePopupPage = `<!doctype html><title>popup</title>
<body><p id="mark">popup</p>
<script>window.open('` + FixtureQuietRoute + `', '_blank');</script></body>`

	FixtureHeavyPagePrefix = `<!doctype html><title>heavy</title>
<body><p id="mark">heavy</p>`

	FixtureHeavyPageSuffix = `</body>`

	FixtureHeavyRowFormat = "<div class=\"row\">row %d padding text</div>"
)
