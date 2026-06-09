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

<div class="wishlist-action">
    <button class="btn-wishlist" onclick="addToWishlist('{{.Country.Name}}')">
        Add to Wishlist
    </button>
    <div id="wishlist-feedback"></div>
</div>

<div class="dest-body">
    <div class="weather-card">
        <h2>Travel weather</h2>
        {{if .Weather}}
        <div class="weather-info">
            <img src="{{.Weather.Icon}}" alt="weather icon">
            <div>
                <span class="temp">{{.Weather.TempC}}°C</span>
                <span class="cond">{{.Weather.Condition}}</span>
            </div>
            <div class="weather-details">
                <span>💧 {{.Weather.Humidity}}%</span>
                <span>💨 {{.Weather.WindKph}} km/h</span>
            </div>
        </div>
        {{else}}
        <p class="weather-note">
            Weather Data is not available for this time.
        </p>
        {{end}}
    </div>

    <div class="attractions-card">
        <h2>Attractions &amp; landmarks</h2>
        <div class="attraction-list">
            {{range .Attractions}}
            <div class="attraction-item">
                <span class="attr-name">{{.Name}}</span>
                <span class="attr-kind">{{.Kinds}}</span>
            </div>
            {{else}}
            <p style="color:#888;font-size:0.9rem">No attractions found for this location.</p>
            {{end}}
        </div>
    </div>
</div>

<script src="/static/js/destination.js"></script>
