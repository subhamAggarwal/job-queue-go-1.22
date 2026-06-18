# Distributed Job Queue System

Welcome to the Advanced Go 1.22 Backend Challenge!

Your goal is to build a high-performance, concurrent Distributed Job Queue System (like a mini-Celery or RabbitMQ) using **Go 1.22**. 
You must use the standard `net/http` library. You are provided with `github.com/gorilla/websocket` for the WebSocket layer and `github.com/mattn/go-sqlite3` for the database.

## Requirements

### 1. `POST /jobs`
- Accept a JSON body containing `payload` (string), `priority` (integer 1-10), and `max_retries` (integer).
- Insert a new job into the `jobs` SQLite table with `status = "pending"`.
- Return the newly created job as JSON with a `201 Created` status code.

### 2. `GET /ws` (Worker Connections)
- Upgrade the HTTP connection to a WebSocket connection.
- When a worker connects, the server must automatically dispatch the highest-priority `pending` job to them via a JSON broadcast.
- The worker will process the job and send a WebSocket message back with `{"action": "complete", "job_id": 123}` or `{"action": "fail", "job_id": 123}`.
- If a job completes, update its status to `completed`.
- If a job fails, increment its `retries` counter. If `retries < max_retries`, set status back to `pending` so another worker can pick it up. If it exceeds retries, set status to `failed`.

### 3. Worker Concurrency & Race Conditions
- Multiple workers will connect simultaneously!
- You MUST ensure thread safety using `sync.Mutex` or Go Channels. Two workers cannot be assigned the same `pending` job.

### 4. Heartbeats (Failure Detection)
- Workers will send a `{"action": "ping"}` message every 3 seconds.
- If the server does not receive a ping from a working worker within **5 seconds**, or if the WebSocket connection drops abruptly, the server MUST immediately mark their assigned job as "failed" and handle retry logic just like a standard failure.

## Development Guide

### How to Run Locally

> **⚠️ CRITICAL: Port Binding**
> To be evaluated, your servers MUST bind to `0.0.0.0` (all interfaces) rather than `localhost` or `127.0.0.1`.

This is a full-stack application. You must run both the backend and frontend servers simultaneously in separate terminal tabs.

**1. Start the Backend Server (Port 8080)**
```bash
go run main.go
```

**2. Start the Frontend Server (Port 5173)**
```bash
cd frontend && npm run dev
```

### How to Test

You can execute the visible test suite locally to verify your solution against the visible test cases:
```bash
go test ./...
```

When you click **Submit**, the platform will run an identical evaluation suite against your code using the same command.
