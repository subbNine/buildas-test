# Go CRUD Application

## Setup

1. Clone the repository.
2. Run `docker-compose up --build` to start the application.

The app will be running at `http://localhost:8080`.

## Endpoints

- `POST /register`: Register a new user.
- `POST /login`: Login and get a JWT token.
- `GET /users`: List all registered users (JWT token required).

## Testing

To run the tests:

```bash
go test ./...
