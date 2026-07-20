package constant

const ConnectionIndicatorScript = `
	document.addEventListener('htmx:sseOpen', function() {
		document.getElementById('sse-dot').className = 'status-dot status-active';
	});
	document.addEventListener('htmx:sseError', function() {
		document.getElementById('sse-dot').className = 'status-dot status-disconnected';
	});
`
