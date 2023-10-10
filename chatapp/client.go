package main

import (
	"bufio"
	"fmt"
	"io"
	"net"
	"os"
	"strings"
	"time"
)

const (
 SERVER_HOST = "192.168.1.19"
 SERVER_PORT = "9988"
 SERVER_TYPE = "tcp"
)

var file, err = os.OpenFile("./mylogfile.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)

var senderIP string = "192.168.1.19"
var receiverIp string

func main() {
 connection, err := net.Dial(SERVER_TYPE, SERVER_HOST+":"+SERVER_PORT)
   if err != nil {
   fmt.Println("Error connecting to server:", err.Error())
   os.Exit(1)
 }
   defer connection.Close()

 if err != nil {
   panic(err)
 }

 clientAddr := connection.RemoteAddr().String()  // RemoteAddr gets the remote address i.e. IP of client and converts it into a string



 go func() { serverStart(clientAddr, connection) }()
   fmt.Print("\nHost to connect to: ") // Asks for input to user
   var host string
   fmt.Scanln(&host)  //fmt.Scanln function reads the user's input and assigns it to the host variable

   ip, _, err := net.SplitHostPort(host)  // net.SplitHostPort splits the input into I.P. address and Port number
   if err != nil {
   fmt.Println("Invalid input format. Please provide IP address and port.")
   return
 }

 receiverIp = ip

 prevMessages()

 for {

   scanner := bufio.NewScanner(os.Stdin)  // bufio.NewScanner creates a new scanner that reads from the standard input
   scanner.Scan() // 

   currentTime := time.Now()  // Returns current time

   timeString := currentTime.Format("2006-01-02 15:04:05")  // Formats the current time in "2006-01-02 15:04:05" format

   text := senderIP + " " + receiverIp + " " + "Kaushal : " + timeString + " " + scanner.Text() + "\n"  // SenderIP, receiver IP and time along with 

   _, err := io.WriteString(file, text)  // writes the text string to the file object
   
   if err != nil {
      panic(err)
    }

   SendMessage(timeString+" "+scanner.Text()+"\n", host)  // sends message with time and message
 
 }
}



func prevMessages() {   // Loads previous messages based on ip address

 file, err := os.Open("./mylogfile.txt")  // Opens the file if any error in opening file it prints the error
   if err != nil {
   fmt.Println(err)
   return
 }

 defer file.Close() // Closes file after function is executed

 scanner := bufio.NewScanner(file) // bufio.NewScanner scans the file line by line and stores it in scanner

 target := receiverIp  // recieverIP is stored in target to use later

 for scanner.Scan() {
 line := scanner.Text() 
 words := strings.Fields(line) // each line is split into words 

 if len(words) >= 2 && words[1] == target {  // checks if atleast 2 words present and if 2nd word matches the target i.e. receiver IP, if both conditions are true then it extracts all words from 3rd word
   arr := append(words[2:])  // extracted words are appended to array named as "arr"
   finalOutput := strings.Join(arr, " ") // words are joined into single string with space
   fmt.Println(strings.Trim(finalOutput, "[]"))   // trims any "[]" brackets
 }
 }

 if err := scanner.Err(); err != nil {  // If any error , it is printed 
 fmt.Println(err)
 return
 }
}

func serverStart(clientAddr string, connection net.Conn) {  // Function starts server , takes two arguments ip address of client and connection object

 PORT := ":" + os.Args[1]  

 l, err := net.Listen("tcp", PORT)   // listens for incoming tcp connection on the port number

 if err != nil {
   fmt.Println(err)
   return
 }
 
 defer l.Close()

 for {  // for loop continously accepts incoming connections 
   c, err := l.Accept()
   if err != nil {
   fmt.Println(err)
   return
 }

   go HandleNewMsg(c, connection)  // For each accepted connection, the HandleNewMsg function is called concurrently

 }
}

func HandleNewMsg(c net.Conn, connection net.Conn) {  // Takes 2 arg, "c" represents the client connection and "connection" represents the server connection, net.Conn is 

 clientAddr := c.RemoteAddr().String()  // c.RemoteAddr() gets clients address as a string 

 receiverIp = strings.Split(clientAddr, ":")[0]  // clients address is split until ":" and stores it as reciever IP 

 for {
   netData, err := bufio.NewReader(c).ReadString('\n')  // clients message is stored in netData variable
   
   // bufio.NewReader(c) reads data from the client connection, ReadString('\n') reads until it encounters newline character
   
   if err != nil {
     fmt.Println(err)
     return
   }

 msg := strings.TrimSpace(string(netData))  // trims clients message "strings.TrimSpace" removes any white space

 println(msg)  // prints msg

 text := senderIP + " " + receiverIp + " " + msg + "\n"  // creates new string concatenating sender IP, reciever IP and msg

 io.WriteString(file, text)  // writes the message to file using "io.WriteString"

 if err != nil {
   panic(err)
 }

 c.Close()  // closes the client connection

 break

 }
}

func SendMessage(message, host string) {  // takes 2 arg, message and host i.e. ip and port of other user

 CONNECT := host  // assigns the host value to 'CONNECT' variable
 c, err := net.Dial("tcp", CONNECT)  // creates TCP connection with other user
   if err != nil {
     fmt.Println(err)
     return
 }

 c.Write([]byte("Shubham :" + " " + message + "\n"))  // writes message recieved from client 
 bufio.NewReader(c).ReadString('\n')  // reads message from client and reads until newline character

}