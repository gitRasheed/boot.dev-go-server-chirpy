# Chirpy Server

Go microservice for user accounts, authentication, and chirp 🐤 posting. Supports JWT auth, refresh tokens, Chirpy Red membership, and full CRUD.

## Features

* User signup/login/update
* JWT access + refresh tokens
* Chirp create/read/delete
* Profanity filter
* Chirpy Red upgrades via Polka webhook
* Query params on chirps:

  * `author_id`
  * `sort=asc|desc`

## Endpoints

| Method | Path                | Description                |
| ------ | ------------------- | -------------------------- |
| POST   | /api/users          | Create user                |
| PUT    | /api/users          | Update user (auth)         |
| POST   | /api/login          | Login, return tokens       |
| POST   | /api/refresh        | Refresh access token       |
| POST   | /api/revoke         | Revoke refresh token       |
| POST   | /api/chirps         | Create chirp (auth)        |
| GET    | /api/chirps         | List chirps                |
| GET    | /api/chirps/{id}    | Get chirp                  |
| DELETE | /api/chirps/{id}    | Delete chirp (author only) |
| POST   | /api/polka/webhooks | Polka webhook              |

## Polka Webhook Format

```json
{
  "event": "user.upgraded",
  "data": { "user_id": "UUID" }
}
```

Auth header:

```text
Authorization: ApiKey <POLKA_KEY>
```

## Tech

* Go 1.21+
* PostgreSQL
* sqlc
* JWT

## Run

```text
go run .
```

`.env`:

```.env
DB_URL=postgres://...
JWT_SECRET=...
POLKA_KEY=...
PLATFORM=dev
```
