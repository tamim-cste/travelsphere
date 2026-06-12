document.addEventListener('DOMContentLoaded', function () {
    loadAllCountries();
});

function loadAllCountries() {
    showLoading();
    fetch('/api/countries?search=&region=')
        .then(res => res.json())
        .then(data => {
            if (!data.success || !data.data) {
                showError('Could not load countries.');
                return;
            }
            renderCountries(data.data);
        })
        .catch(() => showError('Something went wrong.'));
}

function filterCountries() {
    const search = document.getElementById('country-search').value;
    const region = document.getElementById('region-filter').value;

    showLoading();

    fetch(`/api/countries?search=${encodeURIComponent(search)}&region=${encodeURIComponent(region)}`)
        .then(res => res.json())
        .then(data => {
            if (!data.success || !data.data) {
                showError('No countries found.');
                return;
            }
            renderCountries(data.data);
        })
        .catch(() => showError('Something went wrong.'));
}

function renderCountries(countries) {
    if (!countries || countries.length === 0) {
        document.getElementById('country-results').innerHTML =
            '<p style="padding:2rem;color:#888">No results found.</p>';
        return;
    }

    const html = countries.map(c => `
        <a href="/countries/${c.slug}" class="country-card-detail">
            <div class="card-flag">
                <img src="${c.flag}" alt="${c.name}" onerror="this.style.display='none'">
            </div>
            <div class="card-info">
                <h3>${c.name}</h3>
                <p><strong>Capital:</strong> ${c.capital || '—'}</p>
                <p><strong>Population:</strong> ${formatPop(c.population)}</p>
                <p><strong>Currency:</strong> ${c.currency || '—'}</p>
                <p><strong>Languages:</strong> ${c.languages || '—'}</p>
            </div>
        </a>
    `).join('');

    document.getElementById('country-results').innerHTML =
        `<div class="country-grid-list">${html}</div>`;
}

function showLoading() {
    document.getElementById('country-results').innerHTML =
        '<p style="padding:2rem;color:#888">Loading countries...</p>';
}

function showError(msg) {
    document.getElementById('country-results').innerHTML =
        `<p style="padding:2rem;color:#e74c3c">${msg}</p>`;
}

function formatPop(n) {
    if (!n) return '—';
    if (n >= 1e9) return (n / 1e9).toFixed(1) + 'B';
    if (n >= 1e6) return (n / 1e6).toFixed(1) + 'M';
    if (n >= 1e3) return (n / 1e3).toFixed(1) + 'K';
    return n;
}