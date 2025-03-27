package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func main() {
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

	stdinReader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("You: ")
		text, _ := stdinReader.ReadString('\n')
		fmt.Fprint(conn, text)
	}
}

