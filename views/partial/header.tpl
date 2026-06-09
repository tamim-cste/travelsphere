<header>
    <nav>
        <a href="/">TravelSphere</a>
        <a href="/countries" {{if eq .CurrentPath "/countries"}}class="active"{{end}}>
            Explore
        </a>
        <a href="/wishlist" {{if eq .CurrentPath "/wishlist"}}class="active"{{end}}>
            Wishlist
        </a>
        <a href="/dashboard" {{if eq .CurrentPath "/dashboard"}}class="active"{{end}}>
            Dashboard
        </a>
    </nav>
</header>