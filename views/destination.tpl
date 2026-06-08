<!-- Country Header Card -->
<div class="dest-header">
    <div class="dest-flag">
        <img src="{{.Country.FlagURL}}" alt="{{.Country.Name}}">
    </div>
    <div class="dest-info">
        <span class="region-badge">{{.Country.Region}}</span>
        <h1>{{.Country.Name}}</h1>
        <p class="subregion">{{.Country.Subregion}}</p>
        
        <div class="dest-meta">
            <div class="meta-item">
                <span class="meta-label">CAPITAL</span>
                <span>{{.Country.Capital}}</span>
            </div>
            <div class="meta-item">
                <span class="meta-label">POPULATION</span>
                <span>{{.Country.Population}}</span>
            </div>
            <div class="meta-item">
                <span class="meta-label">REGION</span>
                <span>{{.Country.Region}}</span>
            </div>
            <div class="meta-item">
                <span class="meta-label">CURRENCY</span>
                <span>{{.Country.Currency}}</span>
            </div>
            <div class="meta-item">
                <span class="meta-label">LANGUAGES</span>
                <span>{{.Country.Languages}}</span>
            </div>
        </div>
    </div>
</div>

<!-- Add to Wishlist Button -->
<div class="wishlist-action">
    <button class="btn-wishlist" 
            onclick="addToWishlist('{{.Country.Name}}')">
        Add to Wishlist
    </button>
    <div id="wishlist-feedback"></div>
</div>

<!-- Two column layout -->
<div class="dest-body">
    <!-- Weather (left) -->
    <div class="weather-card">
        <h2>Travel weather</h2>
        <p class="weather-note">
            Weather data is optional. Add 
            <code>WEATHER_API_KEY</code> 
            to your <code>.env</code> file to enable live conditions.
        </p>
    </div>

    <!-- Attractions (right) -->
    <div class="attractions-card">
        <h2>Attractions & landmarks</h2>
        <div class="attraction-list">
            {{range .Attractions}}
            <div class="attraction-item">
                <span class="attr-name">{{.Name}}</span>
                <span class="attr-kind">{{.Kinds}}</span>
            </div>
            {{else}}
            <p>No attractions found</p>
            {{end}}
        </div>
    </div>
</div>

<script src="/static/js/destination.js"></script>