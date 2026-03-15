# Chirpy

Chirpy is a RESTful API server for a social media platform, similar to Twitter, built in Go. Users can create accounts, post short messages ("chirps"), and manage their content through a clean JSON API.

## Features

- **User management** - Create accounts, log in, and update profiles
- **JWT authentication** - Secure endpoints with access tokens and refresh tokens
- **Chirps** - Create, read, and delete short-form posts
- **Sorting & filtering** - Query chirps by author and sort by date (ascending or descending)
- **Webhooks** - Integrate with external services (Polka) for premium upgrades
- **Admin dashboard** - Track server metrics and manage state

## Tech Stack

- **Go** standard library (`net/http`) for routing and HTTP handling
- **PostgreSQL** for data persistence
- **sqlc** for type-safe SQL query generation
- **goose** for database migrations

## Getting Started

### Prerequisites

- [Go](https://go.dev/) 1.22+
- [PostgreSQL](https://www.postgresql.org/)

### Installation

1. Clone the repository:
   ```sh
   git clone https://github.com/dylansawicki15/chirpy.git
   cd chirpy
   ```

2. Install dependencies:
   ```sh
   go mod download
   ```

3. Set up the database:
   ```sh
   goose -dir sql/schema postgres "your_connection_string" up
   ```

4. Create a `.env` file in the project root:
   ```
   DB_URL=postgres://user:password@localhost:5432/chirpy?sslmode=disable
   PLATFORM=dev
   SECRET=your-jwt-secret
   POLKA_KEY=your-polka-api-key
   ```

5. Run the server:
   ```sh
   go build -o chirpy && ./chirpy
   ```

   The server starts on `http://localhost:8080`.

## API Endpoints

| Method | Path | Auth | Description |
|--------|------|------|-------------|
| GET | `/api/healthz` | No | Health check |
| POST | `/api/users` | No | Create a user |
| POST | `/api/login` | No | Log in |
| POST | `/api/refresh` | Refresh token | Refresh access token |
| POST | `/api/revoke` | Refresh token | Revoke refresh token |
| PUT | `/api/users` | Access token | Update email/password |
| POST | `/api/chirps` | Access token | Create a chirp |
| GET | `/api/chirps` | No | List chirps (supports `author_id` and `sort` query params) |
| GET | `/api/chirps/{chirpID}` | No | Get a single chirp |
| DELETE | `/api/chirps/{chirpID}` | Access token | Delete your own chirp |
| POST | `/api/polka/webhooks` | API key | Polka webhook for premium upgrades |
