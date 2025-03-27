package main

import (
	"bufio"
	"fmt"
	"io"
	"net"
	"sync"
)

type Hub struct {
	clients map[net.Conn]bool
	mutex   sync.Mutex
	broadcast chan string
}

func NewHub() *Hub {
	return &Hub{
		clients: make(map[net.Conn]bool),
		broadcast: make(chan string),
	}
}

func (h *Hub) Run() {
	for msg := range h.broadcast {
		h.mutex.Lock()
		for conn := range h.clients {
			fmt.Fprintln(conn, msg)
		}
		h.mutex.Unlock()
	}
}

func handleClient(conn net.Conn, hub *Hub) {
	hub.mutex.Lock()
	hub.clients[conn] = true
	hub.mutex.Unlock()

	defer func() {
		hub.mutex.Lock()
		delete(hub.clients, conn)
		hub.mutex.Unlock()
		conn.Close()
	}()

	reader := bufio.NewReader(conn)
	for {
		msg, err := reader.ReadString('\n')
		if err == io.EOF {
			break
		}
		if err != nil {
			continue
		}
		hub.broadcast <- msg
	}
}

