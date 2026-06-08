<section class="hero">
    <h1>Discover Your Next Adventure</h1>
    
    <div class="search-box">
        <input type="text" id="search-input" 
               placeholder="Search destinations...">
        <button onclick="searchDestinations()">Search</button>
    </div>

    <div id="search-suggestions"></div>
</section>

<section class="featured">
    <h2>Featured Destinations</h2>
    <p>Total: {{len .FeaturedCountries}}</p>
    
    <div class="country-grid">
        {{range .FeaturedCountries}}
        <a href="/countries/{{.Slug}}" class="country-card">
            <img src="{{.FlagURL}}" alt="{{.Name}} flag">
            <h3>{{.Name}}</h3>
            <p>{{.Capital}}</p>
        </a>
        {{end}}
    </div>
</section>

<script src="/static/js/home.js"></script>