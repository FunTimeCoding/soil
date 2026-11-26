package constant

const FixturePage = `<title>Probe</title>
<button onclick="document.getElementById('out').textContent='clicked'">Press</button>
<label>Name <input id=i></label>
<p id=out>idle</p>`

const FixtureLinkPage = `<title>Home</title>
<a href="/next">Next</a>
<a href="/slow">Slow</a>
<button onclick="setTimeout(() => location.href = '/next', 50)">Later</button>`

const FixtureNextPage = `<title>Next</title><p>arrived</p>`
