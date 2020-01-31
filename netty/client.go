package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func main() {

	fmt.Println("start client...")

	conn, _ := net.Dial("tcp", "127.0.0.1:8090")

	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(conn)

	for {
		// read in input from stdin
		fmt.Print("Text to send: ")
		text, _ := reader.ReadString('\n')
		// send to socket
		//fmt.Fprintf(conn, text + "\n")
		writer.WriteString(text)
		writer.Flush()
		// listen for reply
		message, _ := bufio.NewReader(conn).ReadString('\n')
		fmt.Print("Message from server: "+message + " \n")
	}

}
