<section class="hero">
    <h1>Discover your next destination</h1>
    <p>Search countries, explore attractions, and curate your personal travel wishlist.</p>
    
    <div class="search-wrapper">
        <label>WHERE TO NEXT?</label>
        <input type="text" id="search-input" 
               placeholder="Search destinations..."
               oninput="searchDestinations()">
        <div id="search-suggestions"></div>
    </div>
</section>

<div class="section">
    <h2>Featured destinations</h2>
    <div class="country-grid">
        {{range .FeaturedCountries}}
        <a href="/countries/{{.Slug}}" class="country-card">
            <img src="{{.FlagURL}}" alt="{{.Name}}">
            <div class="country-card-body">
                <h3>{{.Name}}</h3>
                <p>{{.Capital}} · {{.Region}}</p>
            </div>
        </a>
        {{end}}
    </div>
</div>

<div class="section">
    <h2>Popular attractions</h2>
    <div class="attraction-list" id="attractions-list">
        <div class="attraction-item">
            <span>Eiffel Tower</span>
            <span class="kind">architecture/historic</span>
        </div>
        <div class="attraction-item">
            <span>Grand Canyon</span>
            <span class="kind">nature</span>
        </div>
        <div class="attraction-item">
            <span>Sydney Opera House</span>
            <span class="kind">architecture,theatre</span>
        </div>
        <div class="attraction-item">
            <span>Colosseum</span>
            <span class="kind">historic,architecture</span>
        </div>
    </div>
</div>

<script src="/static/js/home.js"></script>