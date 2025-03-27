package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"

	"github.com/james-w/messaging/shared"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Enter your username: ")
	username, _ := reader.ReadString('\n')
	username = strings.TrimSpace(username)

	if len(username) == 0 {
		fmt.Println("Username cannot be empty.")
		return
	}

	conn, err := net.Dial("tcp", "localhost:9000")
	if err != nil {
		fmt.Println("Unable to connect to server:", err)
		return
	}
	defer conn.Close()

	go func() {
		serverReader := bufio.NewReader(conn)
		for {
			raw, err := serverReader.ReadString('\n')
			if err != nil {
				fmt.Println(err)
				break
			}
			msg, err := shared.Decode([]byte(raw))
			if err != nil {
				break
			}
			fmt.Print(fmt.Sprintf(">> [%s] %s", msg.Username, msg.Text))
		}
	}()

	for {
		fmt.Print("You: ")
		text, _ := reader.ReadString('\n')
		text = strings.TrimSpace(text)

		if text == "" {
			continue
		}

		msg := shared.Message{
			Username: username,
			Text:     text,
		}

		data, err := shared.Encode(msg)
		if err != nil {
			fmt.Println("Error encoding message:", err)
			continue
		}

		fmt.Fprintln(conn, string(data))
	}
}

