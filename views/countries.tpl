<div class="page-header">
    <h1>Country Explorer</h1>
    <p>Browse every destination on first load. Search and filter update only the results below — no full page reload.</p>
</div>

<div class="search-bar-box">
    <div class="search-field">
        <label>SEARCH</label>
        <input type="text" id="country-search" 
               placeholder="Country or capital..."
               oninput="filterCountries()">
    </div>
    <div class="region-field">
        <label>REGION</label>
        <select id="region-filter" onchange="filterCountries()">
            <option value="">All regions</option>
            <option value="Africa">Africa</option>
            <option value="Americas">Americas</option>
            <option value="Asia">Asia</option>
            <option value="Europe">Europe</option>
            <option value="Oceania">Oceania</option>
        </select>
    </div>
</div>

<!-- This div will be replaced with AJAX -->
<div id="country-results">
    {{template "partial/country_grid.tpl" .}}
</div>

<script src="/static/js/countries.js"></script>