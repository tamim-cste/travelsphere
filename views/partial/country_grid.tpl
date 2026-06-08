<div class="country-grid-list">
    {{range .Countries}}
    <a href="/countries/{{.Slug}}" class="country-card-detail">
        <div class="card-flag">
            <img src="{{.FlagURL}}" alt="{{.Name}}">
        </div>
        <div class="card-info">
            <h3>{{.Name}}</h3>
            <p><strong>Capital:</strong> {{.Capital}}</p>
            <p><strong>Population:</strong> {{.Population}}</p>
            <p><strong>Currency:</strong> {{.Currency}}</p>
            <p><strong>Languages:</strong> {{.Languages}}</p>
        </div>
    </a>
    {{end}}
</div>