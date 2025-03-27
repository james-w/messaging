# 📨 Go Messaging Service

A simple TCP-based messaging service built in Go, consisting of a server that broadcasts incoming messages to all connected clients. It serves as a minimal yet extensible foundation for exploring networking, concurrency, and distributed system concepts.

---

## 🏗 Project Structure

```
messaging-service-go/
├── server/
│   ├── main.go            # Server entry point
│   └── broadcast.go       # Hub for broadcasting to clients
├── client/
│   └── main.go            # Console-based chat client
├── shared/
│   └── protocol.go        # Placeholder for future message struct/protocol
└── go.mod
```

---

## 🚀 Getting Started

### Prerequisites
- Go 1.21+
- Terminal (for running multiple clients)

### Run the Server

```bash
go run ./server/*
```

### Run One or More Clients

```bash
go run ./client/*
```

Type into any client to send a message. All connected clients will receive it.

---

## ⚙ Features

- [x] Multiple clients connect over TCP
- [x] Messages are broadcast to all other clients
- [x] Concurrent client handling using goroutines
- [x] Simple stdin/stdout-based CLI

---

## 🧪 Example

```
# Terminal 1 (Server)
Server listening on :9000

# Terminal 2 (Client A)
You: Hello world!
>> Hello world!

# Terminal 3 (Client B)
>> Hello world!
You: 👋
>> 👋
```

---

## 📌 Opportunities for Extension

- Add client identity (usernames)
- Structured messages with timestamps (see `shared/protocol.go`)
- Support private messages or chat rooms
- Add WebSocket support for browser-based client
- Logging, persistence, or message replay
- Basic authentication or client rate limiting

