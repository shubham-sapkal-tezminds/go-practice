package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

const (
 SERVER_TYPE = "tcp"
)

func main() {

    fmt.Println("Server Is Running....")

    SERVER_HOST := ""
    SERVER_PORT := ""

    fmt.Println("Enter server IP address :- ")

    fmt.Scanln(&SERVER_HOST)

    fmt.Println("Enter port number :- ")

    fmt.Scanln(&SERVER_PORT)

    fmt.Printf("Value of ip address is  %v and port number is %v :- ", SERVER_HOST, SERVER_PORT)

    server, err := net.Listen(SERVER_TYPE, SERVER_HOST+":"+SERVER_PORT)  // Listen function of the net package in Go creates servers and listens for incoming connections on a local network address
 
    checkNilErr(err)  // checks if any error
 
    
    
    defer server.Close()

    fmt.Println("Listening on " + SERVER_HOST + ":" + SERVER_PORT)

    fmt.Println("Waiting for Client...")

    for {
      conn, err := server.Accept()  // Accept method of a Listener interface and accepts incoming client connections

      checkNilErr(err)

      fmt.Println("Client Connected")

      go processClient(conn)

      continue

    }
}

func processClient(conn net.Conn) {

    buffer := make([]byte, 1024)  // buffer stores data in bytes by creating slice of 1024 bytes

    mLen, err := conn.Read(buffer) // it reads data from conn and stores it in mlen

    checkNilErr(err)

    fmt.Println("Received: ", string(buffer[:mLen]))   // converts to string from bytes and removes whitespace after the message to print it

    fmt.Print("Enter message to send: ")

    reader := bufio.NewReader(os.Stdin)  // it reads message and stores it in reader

    message, _ := reader.ReadString('\n') // message is converted to string

    _, err = conn.Write([]byte(message)) // conn.Write method is used to write data to a network connection. It takes a byte slice as its argument, and returns the number of bytes written or an error.

    checkNilErr(err)

   
   if message == "bye" {
     fmt.Println("exit message",message)
     conn.Close()

   }
}


func checkNilErr(err error) {
	if err != nil {
		fmt.Println("Error reading: ", err.Error())
		os.Exit(1)
	}
}