#!/usr/bin/env bash
set -e

echo "Starting Go 1.22 Job Queue Backend..."
go run main.go &

echo "Starting React Frontend on port 5173..."
cd frontend
if [ ! -d "node_modules" ]; then
  npm install
fi

# Vite explicitly bound to 0.0.0.0 on 5173
npm run dev -- --host 0.0.0.0 --port 5173
