# Message Board

Message board Go application. Multiple database types are available.
A database is selected by specifying the `DB_TYPE` environment variable.

| Database           | DB_TYPE Variable | Notes                                                                                        |
| ------------------ | ---------------- | -------------------------------------------------------------------------------------------- |
| Postgres (default) | postgres         | Database connection variables include: `DB_HOST`, `DB_PORT`, `DB_NAME`, `DB_USER`, `DB_PASS` |
| SQLite             | sqlite           | Optionally specify SQLite file in `SQLITE_FILE` environment variable.                        |
| SQLite In-Memory   | sqlite-memory    | Temporary in-memory SQLite database. Contents are deleted on service shutdown.               |

## Run Locally

```
cd message-board
DB_TYPE=sqlite-memory go run .
```

## Run Docker

```
cd message-board
docker build -t messageboard:latest .
```

## Run Tests

```
cd message-board
go test ./...
```

## Endpoints

The message board service contains the following endpoints:

| Endpoint          | Description           |
| ----------------- | --------------------- |
| `POST /message`   | Create a new message. |
| `GET /message`    | Get a message by ID.  |
| `PATCH /message`  | Update a message.     |
| `DELETE /message` | Delete a message.     |
