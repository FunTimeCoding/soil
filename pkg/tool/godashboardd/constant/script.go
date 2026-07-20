package constant

const BeaconScript = `
document.addEventListener('click', function(e) {
	var link = e.target.closest('.board-entry-link');
	if (!link || !link.dataset.label) return;
	navigator.sendBeacon('/click', link.dataset.label);
});
`
