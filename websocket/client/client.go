package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

const (
 SERVER_HOST = "192.168.1.41"
 SERVER_PORT = "8080"
 SERVER_TYPE = "tcp"
)

func main() {
 fmt.Println("Client Is Running...")       
 
 conn, err := net.Dial(SERVER_TYPE, SERVER_HOST+":"+SERVER_PORT)  // net.Dial connects to server , taking 2 arguments . 1st :- type of network, 2nd :- address
 
 checkNilErr(err) // checks if any error present
 
 
for {

	fmt.Print("Send Message: ")
	
	reader := bufio.NewReader(os.Stdin) // it reads message and stores it in reader
	message, _ := reader.ReadString('\n') // message is converted to string
	
	_, err = conn.Write([]byte(message)) // conn.Write method is used to write data to a network connection. It takes a byte slice as its argument, and returns the number of bytes written or an error.
	
	checkNilErr(err) // checks if any error present
	
	buffer := make([]byte, 1024) // buffer stores data in bytes by creating slice of 1024 bytes
	
	mLen, err := conn.Read(buffer) // it reads data from conn and stores it in mlen
	
	checkNilErr(err)
	
	fmt.Println(string(buffer[:mLen]))  // converts to string from bytes and removes whitespace after the message

	if string(buffer[:mLen]) == "bye" {
        fmt.Println("exiting client 1")

		break
	}
}

fmt.Println("exiting client 2")
conn.Close() // closes connection at the end



}


func checkNilErr(err error) {
	if err != nil {
		fmt.Println("Error reading: ", err)
		// panic(err)
	}
}