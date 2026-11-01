package constant

const TileFilterScript = `
function filterTiles(input) {
	var lower = input.value.toLowerCase();
	input.nextElementSibling.querySelectorAll('[data-search]').forEach(function(tile) {
		var text = (tile.getAttribute('data-search') || '').toLowerCase();
		tile.style.display = text.includes(lower) ? '' : 'none';
	});
}
`
