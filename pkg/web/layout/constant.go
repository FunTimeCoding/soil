package layout

const SummaryStrip = "summary_strip"
const NotificationRegion = "notifications"
const StatusLine = "status-line"
const StatusSuccess = "success"
const StatusNotice = "notice"
const baseStyle = `
a { text-decoration: none; }
a:hover { text-decoration: none; color: var(--pico-primary-hover); }
`
const notificationStyle = `
#notifications:empty { display: none; }
#notifications { margin-bottom: 1rem; }
.notification-error {
	display: flex;
	align-items: baseline;
	gap: 0.75rem;
	padding: 0.75rem 1rem;
	margin-bottom: 0.5rem;
	border: 1px solid var(--pico-del-color);
	border-left-width: 4px;
	border-radius: var(--pico-border-radius);
	background: var(--pico-card-background-color);
}
.notification-message { flex: 1; }
.notification-event {
	font-family: monospace;
	color: var(--pico-muted-color);
}
.notification-close {
	cursor: pointer;
	background: none;
	border: none;
	color: var(--pico-muted-color);
	padding: 0;
	width: auto;
	line-height: 1;
}
#status-line:empty { display: none; }
.notification-status {
	display: flex;
	align-items: baseline;
	gap: 0.75rem;
	padding: 0.75rem 1rem;
	margin-bottom: 0.5rem;
	border: 1px solid var(--pico-muted-color);
	border-left-width: 4px;
	border-radius: var(--pico-border-radius);
	background: var(--pico-card-background-color);
}
.notification-status.notification-success {
	border-color: var(--pico-ins-color);
}
`
const notificationScript = `
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
const paletteStyle = `
.palette-dialog {
	position: fixed;
	top: 15vh;
	left: 50%;
	transform: translateX(-50%);
	min-width: auto;
	min-height: auto;
	width: min(600px, 90vw);
	max-height: 70vh;
	margin: 0;
	border: 1px solid var(--pico-muted-border-color);
	border-radius: var(--pico-border-radius);
	padding: 0;
	overflow: hidden;
	background: var(--pico-card-background-color);
	color: var(--pico-color);
	box-shadow: 0 8px 32px rgba(0, 0, 0, 0.4);
	display: flex;
	flex-direction: column;
}
.palette-dialog::backdrop {
	background: rgba(0, 0, 0, 0.5);
	-webkit-backdrop-filter: none;
	backdrop-filter: none;
}
.palette-body {
	display: flex;
	flex-direction: column;
	max-height: 70vh;
}
.palette-input {
	width: 100%;
	padding: 0.75rem 1rem;
	font-size: 1rem;
	border: none;
	border-bottom: 1px solid var(--pico-muted-border-color);
	background: transparent;
	color: var(--pico-color);
	outline: none;
	margin: 0;
}
#palette-results {
	overflow-y: auto;
	max-height: calc(70vh - 3rem);
}
.palette-item {
	display: flex;
	align-items: center;
	gap: 0.75rem;
	padding: 0.6rem 1rem;
	text-decoration: none;
	color: var(--pico-color);
	cursor: pointer;
}
.palette-item:hover,
.palette-item.active {
	background: var(--pico-primary-focus);
}
.palette-label {
	flex: 1;
	font-size: 0.95rem;
}
.palette-label strong {
	color: var(--pico-primary);
}
.palette-description {
	font-size: 0.8rem;
	color: var(--pico-muted-color);
}
.palette-category {
	font-size: 0.75rem;
	padding: 0.1rem 0.4rem;
	border-radius: 4px;
	background: var(--pico-muted-border-color);
	color: var(--pico-muted-color);
	text-transform: uppercase;
}
.palette-empty {
	padding: 1.5rem;
	text-align: center;
	color: var(--pico-muted-color);
}
`

const paletteScript = `
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
