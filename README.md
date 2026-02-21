# Fintech Wallet Backend

A backend service written in Go that simulates core fintech wallet operations such as user creation, wallet management, and secure money transfers.

This project focuses on backend architecture, transactional safety, and clean service design rather than UI.

---

## Tech Stack

- Go
- Gin (HTTP framework)
- PostgreSQL
- Docker
- REST APIs

---

## Features

- User management APIs
- Wallet creation and balance handling
- Atomic money transfers using database transactions
- Row-level locking to prevent race conditions
- Clean layered architecture (handler → service → repository)
- Dockerized PostgreSQL setup

---

## Project Structure

cmd/server - API entrypoint
internal/user - user domain
internal/wallet - wallet domain
db/init - database schema
api-test - HTTP request tests

---

## Running Locally

Start database:

docker compose up -d

Run server:

go run cmd/server/main.go

---

## API Examples

Create user:

POST /users

Transfer funds:

POST /wallets/transfer