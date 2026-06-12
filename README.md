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

## Testing
You can run your server manually in the terminal to test your endpoints with `curl` or Postman:
```bash
go run main.go
```

To run the local visible test cases, click the **Run Code** button in your IDE or run:
```bash
go test -v
```

When you are ready, click **Submit**. Our hidden evaluation suite will spin up your server and simulate dozens of workers, inject race conditions, and abruptly disconnect workers to test your concurrent recovery logic!
