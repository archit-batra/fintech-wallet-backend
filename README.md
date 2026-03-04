# Go Backend Platform

A production-style backend service written in Go that simulates wallet and transfer operations while demonstrating clean architecture, transactional safety, and asynchronous event processing.

This project focuses on backend system design, reliability, and infrastructure patterns rather than UI.

---

## Architecture Overview

### API Service
- REST APIs built using **Gin**
- **PostgreSQL** for transactional data
- **Redis** for event queueing

### Worker Service
- Consumes events from Redis
- Persists audit logs asynchronously
- Supports graceful shutdown

### Concepts Demonstrated
- Atomic database transactions
- Row-level locking for concurrency safety
- Event-driven architecture
- Eventual consistency
- Background job processing

---

## Tech Stack

- Go
- Gin (HTTP framework)
- PostgreSQL
- Redis
- Docker
- REST APIs

---

## Key Features

- User management APIs
- Wallet creation and balance updates
- Atomic money transfers using DB transactions
- Row-level locking to prevent race conditions
- Redis-based async event queue
- Background worker for event processing
- Audit logging via eventual consistency
- Health checks for DB and Redis
- Graceful shutdown support

---

## Project Structure

cmd/server → API entrypoint

cmd/worker → Background worker service

internal/user → User domain

internal/wallet → Wallet domain

internal/audit → Audit logging

internal/events → Event models

internal/infra → Infrastructure (Redis client)

db/init → Database schema

api-test → HTTP request examples

---

## Running Locally

Start infrastructure:
docker compose up -d

Run API server:
go run cmd/server/main.go

Run worker:
go run cmd/worker/main.go

---

## Example Endpoints

Create user
POST /users

Transfer funds
POST /wallets/transfer

Health check
GET /health