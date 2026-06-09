<header>
    <nav>
        <a href="/" class="nav-brand">TravelSphere</a>
        <a href="/" {{if eq .CurrentPath "/"}}class="active"{{end}}>Home</a>
        <a href="/countries" {{if eq .CurrentPath "/countries"}}class="active"{{end}}>Countries</a>
        <a href="/wishlist" {{if eq .CurrentPath "/wishlist"}}class="active"{{end}}>Wishlist</a>
        <a href="/dashboard" {{if eq .CurrentPath "/dashboard"}}class="active"{{end}}>Dashboard</a>
        <div class="nav-right">
            {{if .LoggedIn}}
            <span class="nav-greeting">Hi, {{.Username}}</span>
            <a href="/logout" class="btn-nav">Logout</a>
            {{else}}
            <a href="/login" class="btn-nav">Login</a>
            {{end}}
        </div>
    </nav>
</header>
