package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter your username: ")
	username, _ := reader.ReadString('\n')
	username = username[:len(username)-1] // strip newline

	conn, err := net.Dial("tcp", "localhost:9000")
	if err != nil {
		fmt.Println("Unable to connect to server:", err)
		return
	}
	defer conn.Close()

	go func() {
		serverReader := bufio.NewReader(conn)
		for {
			msg, err := serverReader.ReadString('\n')
			if err != nil {
				break
			}
			fmt.Print(">> " + msg)
		}
	}()

	// reused `reader` above
	for {
		fmt.Print("You: ")
		text, _ := reader.ReadString('\n')
		formatted := fmt.Sprintf("[%s]: %s", username, text)
		fmt.Fprint(conn, formatted)
	}
}

