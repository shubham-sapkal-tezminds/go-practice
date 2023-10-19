package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"net"
	"os"
	"regexp"
	"strings"
	"time"

	"github.com/fatih/color"
	"gitlab.com/david_mbuvi/go_asterisks"
)

const (
 SERVER_HOST = "192.168.1.41"
 SERVER_PORT = "9988"
 SERVER_TYPE = "tcp"
)

var receiverUser string

var username string
var password string
var receiverIp string
var ipAddress string
var host string

func main() {
 fmt.Println()
 // c := color.New(color.FgMagenta)
 color.Cyan("1. Login with same device")
 color.Cyan("2. Login with new device")
 color.Cyan("3. Signup")
 fmt.Println()
 fmt.Print("Enter the option number : ")

 var choice int
 _, err := fmt.Scanln(&choice)
 CheckError(err)

 switch choice {
 case 1:
 Loginwithsame()
 case 2:
 LoginNewUser()
 case 3:
 SignupUser()
 default:
 color.Red("Invalid Choice! Enter Correct Choice")
 }
 fmt.Println()
}

var c = color.New(color.FgYellow)

func Loginwithsame() {
 fmt.Println()
 fmt.Println("Login Details")

 for {
 c.Print("Enter your username : ")
 fmt.Scanln(&username)

 c.Print("Enter your password : ")
 bytePassword, err := go_asterisks.GetUsersPassword("", true, os.Stdin, os.Stdout)
 if err != nil {
 fmt.Println(err.Error())
 }
 password := string(bytePassword)

 if len(username) > 0 && len(password) > 0 {

 var data = []string{username, password, "Loginwithsame"}

 conn, err := net.Dial(SERVER_TYPE, SERVER_HOST+":"+SERVER_PORT)
 CheckError(err)

 defer conn.Close()

 t, err := json.Marshal(data)

 finalOutput := string(t)

 _, err = conn.Write([]byte(finalOutput))

 buffer := make([]byte, 1024)
 n, err := conn.Read(buffer)
 CheckError(err)

 msg := string(buffer[:n])

 if msg == "Login Successful" {
 // fmt.Println("Checking : ", msg)
 handleChats(msg, conn)

 break

 } else {
 fmt.Println()
 color.Red(msg)
 fmt.Println()
 continue
 }

 } else {
 fmt.Println("Fill all the details")
 }
 }
}

func LoginNewUser() {
 fmt.Println()
 fmt.Println("Login Details")

 for {
 c.Print("Enter your username : ")
 fmt.Scanln(&username)

 c.Print("Enter your password : ")
 bytePassword, err := go_asterisks.GetUsersPassword("", true, os.Stdin, os.Stdout)
 if err != nil {
 fmt.Println(err.Error())
 }
 password := string(bytePassword)
 // fmt.Scanln(&password)

 c.Print("Enter IP Address and Port : ")

 fmt.Scanln(&host)

 if len(username) > 0 && len(password) > 0 {

 var data = []string{username, password, "Login", host}

 conn, err := net.Dial(SERVER_TYPE, SERVER_HOST+":"+SERVER_PORT)
 CheckError(err)

 defer conn.Close()

 t, err := json.Marshal(data)

 finalOutput := string(t)

 _, err = conn.Write([]byte(finalOutput))

 buffer := make([]byte, 1024)
 n, err := conn.Read(buffer)
 CheckError(err)

 msg := string(buffer[:n])

 if msg == "Login Successful" {
 // fmt.Println("Checking : ", msg)
 handleChats(msg, conn)

 break

 } else {
 fmt.Println()
 color.Red(msg)
 fmt.Println()
 continue
 }

 } else {
 fmt.Println("Fill all the details")
 }
 }
}

func SignupUser() {
 fmt.Println()
 fmt.Println("Signup Details")

 for {
 c.Print("Enter unique username : ")

 fmt.Scanln(&username)

 c.Print("Enter password : ")
 bytePassword, err := go_asterisks.GetUsersPassword("", true, os.Stdin, os.Stdout)
 if err != nil {
 fmt.Println(err.Error())
 }
 password := string(bytePassword)

 c.Print("Enter IP Address and Port : ")

 fmt.Scanln(&host)

 if len(username) > 0 && len(password) > 0 {

 var data []string

 if isValidPassword(password) {
 data = []string{username, password, "Signup", host}

 conn, err := net.Dial(SERVER_TYPE, SERVER_HOST+":"+SERVER_PORT)
 CheckError(err)

 defer conn.Close()

 t, err := json.Marshal(data)

 finalOutput := string(t)

 _, err = conn.Write([]byte(finalOutput))
 CheckError(err)
 buffer := make([]byte, 1024)
 n, err := conn.Read(buffer)
 CheckError(err)

 msg := string(buffer[:n])

 if msg == "Account created succesfully!!!!" {
 handleChats(msg, conn)
 break
 } else {
 fmt.Println()
 color.Red(msg)
 fmt.Println()
 continue
 }
 } else {
 color.Red("Please enter strong password")
 continue
 }

 } else {
 fmt.Println("Fill all the details")
 }
 }

}

func handleChats(msg string, conn net.Conn) {
 fmt.Println()
 color.Green(msg)
 clientAddr := conn.RemoteAddr().String()

 newConn, err := net.Dial(SERVER_TYPE, SERVER_HOST+":"+SERVER_PORT)
 CheckError(err)

 sendmsg := "/getport:" + username
 newConn.Write([]byte(sendmsg))

 buffer := make([]byte, 1024)
 n, err := newConn.Read(buffer)

 CheckError(err)

 port := strings.TrimSpace(string(buffer[:n]))

 go func() { serverStart(clientAddr, conn, port) }()

 for {
 var color1 = color.New(color.FgBlue)

 var connectWithUser string
 color1.Print("\nEnter user to connect : ")
 fmt.Scanln(&connectWithUser)

 connection, err := net.Dial(SERVER_TYPE, SERVER_HOST+":"+SERVER_PORT)
 CheckError(err)

 message := "/connectwithuser:" + connectWithUser + "#" + username
 connection.Write([]byte(message))

 buffer := make([]byte, 1024)
 n, err := connection.Read(buffer)

 CheckError(err)

 host = strings.TrimSpace(string(buffer[:n]))

 if host == "Username not exist!" {
 fmt.Println()
 color.Red(host)
 continue
 } else {

 for {
 fmt.Println()
 color.Cyan("1. Chat history")
 color.Cyan("2. User status")
 color.Cyan("3. Start conversation")
 fmt.Println()
 fmt.Print("Enter the option number : ")

 var choice int
 _, err := fmt.Scanln(&choice)
 CheckError(err)

 switch choice {
 case 1:
 ChatHistory()
 case 2:
 UserStatus()
 case 3:
 StartConversation()
 default:
 color.Red("Invalid Choice! Enter 1 or 2")
 }
 fmt.Println()

 }

 }
 }

}

func StartConversation() {
 fmt.Println()
 color.Cyan("Start Chatting :")
 fmt.Println()
 UserStatus()
 for {
 scanner := bufio.NewScanner(os.Stdin)
 scanner.Scan()
 mymsg := scanner.Text()

 conn, err := net.Dial(SERVER_TYPE, SERVER_HOST+":"+SERVER_PORT)
 CheckError(err)
 if mymsg == "/exit" {
 msg := "/close:" + username
 conn.Write([]byte(msg))
 os.Exit(0)
 return

 }

 currentTime := time.Now()

 timeString := currentTime.Format("2006-01-02 15:04:05")

 text := username + " : " + timeString + " " + mymsg

 SendMessage(text, host)

 }
}

func UserStatus() {
 conn2, err := net.Dial(SERVER_TYPE, SERVER_HOST+":"+SERVER_PORT)
 CheckError(err)

 msg2 := "/userstatus:" + username
 conn2.Write([]byte(msg2))

 buffer := make([]byte, 1024)
 n, err := conn2.Read(buffer)

 CheckError(err)

 val1 := strings.TrimSpace(string(buffer[:n]))

 fmt.Println()

 arrayOfString := strings.Fields(val1)

 receiverUser = arrayOfString[0] + " went offline"

 if arrayOfString[2] == "online" {
 color.Green(val1)
 } else {
 color.Yellow(val1)
 }
}

func ChatHistory() {
 conn1, err := net.Dial(SERVER_TYPE, SERVER_HOST+":"+SERVER_PORT)
 CheckError(err)

 msg1 := "/previouschat:" + username + "#" + password
 conn1.Write([]byte(msg1))

 buffer1 := make([]byte, 1024)
 n1, err := conn1.Read(buffer1)

 CheckError(err)

 val := strings.TrimSpace(string(buffer1[:n1]))

 if strings.HasPrefix(val, "/allchat:") {
 splitResult := strings.SplitN(val, ":", 2)
 if len(splitResult) == 2 {
 result := strings.TrimSpace(splitResult[1])

 var data1 []string

 _ = json.Unmarshal([]byte(result), &data1)

 fmt.Println()
 color.Cyan("Previous messages :")
 fmt.Println()

 for _, res := range data1 {
 fmt.Println(res)
 }

 } else {
 fmt.Println("Invalid input format")
 }
 }
}

func SendMessage(message, host string) {

 CONNECT := host
 c, err := net.Dial("tcp", CONNECT)
 CheckError(err)

 conn, err := net.Dial(SERVER_TYPE, SERVER_HOST+":"+SERVER_PORT)

 msg := "/sendchat:" + message
 conn.Write([]byte(msg))

 c.Write([]byte(message + "\n"))

 CheckError(err)

}

func HandleNewMsg(c, connection net.Conn) {

 for {
 netData, err := bufio.NewReader(c).ReadString('\n')
 if err != nil {
 if err == io.EOF {
 color.Red(receiverUser)
 c.Write([]byte("/exit\n")) // Send /exit command to server
 return
 }

 fmt.Println(err)
 return
 }

 msg := strings.TrimSpace(string(netData)) //naveen : hey
 var color1 = color.New(color.FgBlue)
 color1.Println(msg)
 }
}

func serverStart(clientAddr string, connection net.Conn, PORT string) {

 l, err := net.Listen("tcp", PORT)

 if err != nil {
 fmt.Println(err)
 return
 }

 defer l.Close()

 for {
 c, err := l.Accept()
 if err != nil {
 fmt.Println(err)
 return
 }
 go HandleNewMsg(c, connection)

 }
}

func CheckError(err error) {
 if err != nil {
 fmt.Println("Error is : ", err)
 return
 }
}

func isValidPassword(password string) bool {
 secure := true
 tests := []string{".{7,}", "[a-z]", "[A-Z]", "[0-9]", "[^\\d\\w]"}
 for _, test := range tests {
 t, err := regexp.MatchString(test, password)
 if err != nil {
 fmt.Printf("Regex error: %v\n", err)
 return false
 }

 if !t {
 secure = false
 break
 }
 }
 return secure
}
