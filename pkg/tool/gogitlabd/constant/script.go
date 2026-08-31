package constant

const ScrollScript = `(() => {
	const pane = document.querySelector('#log-pane pre');

	if (pane) {
		pane.scrollTop = pane.scrollHeight;
	}
})();`
