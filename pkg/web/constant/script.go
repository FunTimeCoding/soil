package constant

const (
	NotificationScript = `
document.body.addEventListener('htmx:responseError', function(e) {
	var region = document.getElementById('notifications');

	if (!region) { return; }

	var xhr = e.detail.xhr;

	if (xhr.getResponseHeader('Notification-Item')) {
		region.insertAdjacentHTML('beforeend', xhr.responseText);

		return;
	}

	var item = document.createElement('div');
	item.className = 'notification-error';
	item.setAttribute('role', 'alert');
	var message = document.createElement('span');
	message.className = 'notification-message';
	var text = (xhr.responseText || '').slice(0, 300);
	message.textContent = 'Request failed (' + xhr.status + '): ' + text;
	var close = document.createElement('button');
	close.className = 'notification-close';
	close.textContent = '×';
	close.addEventListener('click', function() { item.remove(); });
	item.appendChild(message);
	item.appendChild(close);
	region.appendChild(item);
});
`
	PaletteScript = `
(function() {
	var dialog = document.getElementById('palette');
	var input = document.getElementById('palette-input');

	document.addEventListener('keydown', function(e) {
		if ((e.metaKey || e.ctrlKey) && e.key === 'k') {
			e.preventDefault();
			dialog.showModal();
			input.value = '';
			input.focus();
			input.dispatchEvent(new Event('keyup'));
		}
	});

	dialog.addEventListener('keydown', function(e) {
		var items = dialog.querySelectorAll('[data-palette-item]');
		var active = dialog.querySelector('.active');
		var index = Array.prototype.indexOf.call(items, active);

		if (e.key === 'ArrowDown') {
			e.preventDefault();
			if (active) active.classList.remove('active');
			index = (index + 1) % items.length;
			items[index].classList.add('active');
		} else if (e.key === 'ArrowUp') {
			e.preventDefault();
			if (active) active.classList.remove('active');
			index = index <= 0 ? items.length - 1 : index - 1;
			items[index].classList.add('active');
		} else if (e.key === 'Enter' && active) {
			e.preventDefault();
			active.click();
		}
	});

	document.addEventListener('htmx:afterSwap', function(e) {
		if (e.detail.target.id === 'palette-results' ||
			e.detail.target.parentElement &&
			e.detail.target.parentElement.id === 'palette-results') {
			var first = dialog.querySelector('[data-palette-item]');
			if (first) first.classList.add('active');
		}
	});
})();
`
)
