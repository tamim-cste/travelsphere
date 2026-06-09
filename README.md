# TravelSphere

TravelSphere is a Go + Beego web application for exploring countries, viewing destination details, maintaining a wishlist, and checking a simple dashboard summary. The project is designed to run locally with minimal setup and to be containerized with Docker.

## Highlights
- Country search and filtering with AJAX-driven UI
- Destination / attraction data using OpenTripMap
- Optional weather data using WeatherAPI
- Wishlist management with an in-memory store (no database required)
- Docker-ready build and runtime instructions

## Requirements
- Go 1.26+
- Docker (optional, for containerized runs)
- Internet access for external APIs

No database setup is required.

## Environment variables
Create a local `.env` file from the example file before running the app:

```sh
cp .env.example .env
```

The application uses the following variables:

- `OPENTRIPMAP_KEY` (required)
  - Needed for attraction / destination data.
  - If this key is missing, the app should fall back gracefully and show a user-friendly message instead of crashing.
- `WEATHER_API_KEY` (optional)
  - Used for weather lookups.
  - If it is not set, weather-related features are skipped safely.



## Local setup

1. Copy the example environment file:

   ```sh
   cp .env.example .env
   ```

2. Fill in the required values in `.env`:

   ```env
   OPENTRIPMAP_KEY=your_opentripmap_key_here
   WEATHER_API_KEY=your_weatherapi_key_here
   ```

3. Run tests:

   ```sh
   go test ./...
   ```

4. Start the app:

   ```sh
   go run .
   ```

5. Open the app in your browser:

   ```text
   http://localhost:8080
   ```

## Docker setup

Pull from docker hub
```sh
docker pull tamim99/travelsphere
```

### Build locally

```sh
docker build -t travelsphere:local .
```

### Run the container

Use your local environment variables at runtime:

```sh
docker run --rm -p 8080:8080 --env-file .env travelsphere:local
```
