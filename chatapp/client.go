package main

import (
	"bufio"
	"encoding/json"
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

var userNames = map[string]string{
  "naveen": "192.168.1.42:8000",
  "kaushal": "192.168.1.19:6000",
  "vipul" : "192.168.1.40:7000",
}


var file, _ = os.OpenFile("./mylogfile.txt", os.O_CREATE|os.O_WRONLY, 0644)

var senderIP string = "192.168.1.19"
var receiverIp string
var Ip string
type User struct {
    IpAddress string
    UserInfo  UserInfo
}

type UserInfo struct {
    Username string
    Chat     []string
}

var users []User

var data = make(map[string]UserInfo)

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

  

    clientAddr := connection.RemoteAddr().String()

    go func() { serverStart(clientAddr, connection) }()
    fmt.Print("\nHost to connect to: ")
    var host string
    fmt.Scanln(&host)
    fmt.Println("name :",host)


    for key := range userNames { 
        // fmt.Printf("key[%s] value[%s]\n", k, v)
        if host == key {
        var userAddress = userNames[host]

    //    fmt.Println("ip :-",user)
        Ip = userAddress

        fmt.Println(Ip)
    
        } 
    }

    ip, _, err := net.SplitHostPort(Ip)
    if err != nil {
        fmt.Println("Invalid input format. Please provide IP address and port.")
        return
    }

    receiverIp = ip

    prevMessages()

    for {

        scanner := bufio.NewScanner(os.Stdin)
        scanner.Scan()
        mymsg := scanner.Text()

        currentTime := time.Now()

        timeString := currentTime.Format("2006-01-02 15:04:05")

        text := "Shubham : " + timeString + " " + mymsg

        username := strings.Split(text, ":")[0]

        file, err := os.Open("./mylogfile.txt")
        if err != nil {
            fmt.Println(err)
            return
        }
        defer file.Close()

        scan := bufio.NewScanner(file)
        scan.Scan()
        line := scan.Text()

        if err := scan.Err(); err != nil {
            fmt.Println(err)
            return
        }

        var map_1 = make(map[string]UserInfo)

        _ = json.Unmarshal([]byte(line), &map_1)

        if len(map_1) == 0 {
            map_1[receiverIp] = UserInfo{Username: username, Chat: []string{text}}
        }

        mydata := map_1[receiverIp]
        mydata.Chat = append(mydata.Chat, text)
        map_1[receiverIp] = mydata

        t, err := json.Marshal(map_1)
        if err != nil {
            panic(err)
        }

        finalOut := string(t)

        myfile, err := os.OpenFile("./mylogfile.txt", os.O_CREATE|os.O_WRONLY, 0644)
        if err != nil {
            fmt.Println(err)
            return
        }

        _, err = io.WriteString(myfile, finalOut)
        if err != nil {
            fmt.Println(err)
            return
        }

        myfile.Close()

        if err != nil {
            panic(err)
        }

        SendMessage(timeString+" "+scanner.Text()+"\n", Ip)

    }
}

func prevMessages() {

    file, err := os.Open("./mylogfile.txt")

    if err != nil {
        fmt.Println(err)
        return
    }
    defer file.Close()

    scanner := bufio.NewScanner(file)
    scanner.Scan()
    line := scanner.Text()

    if err := scanner.Err(); err != nil {
        fmt.Println(err)
        return
    }

    var map_1 = make(map[string]UserInfo)
    _ = json.Unmarshal([]byte(line), &map_1)

    for _, val := range map_1[receiverIp].Chat {
        fmt.Println(val)
    }
}
func serverStart(clientAddr string, connection net.Conn) {

    PORT := ":" + os.Args[1]

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

func HandleNewMsg(c net.Conn, connection net.Conn) {

    clientAddr := c.RemoteAddr().String()

    receiverIp = strings.Split(clientAddr, ":")[0]

    for {
        netData, err := bufio.NewReader(c).ReadString('\n')

        if err != nil {
            fmt.Println(err)
            return
        }

        msg := strings.TrimSpace(string(netData)) 

        fmt.Println(msg)

        username := strings.Split(msg, ":")[0]

        file, err := os.Open("./mylogfile.txt")
        if err != nil {
            fmt.Println(err)
            return
        }
        defer file.Close()

        scanner := bufio.NewScanner(file)
        scanner.Scan()
        line := scanner.Text()

        if err := scanner.Err(); err != nil {
            fmt.Println(err)
            return
        }

        var map_1 = make(map[string]UserInfo)

        _ = json.Unmarshal([]byte(line), &map_1)

        if len(map_1) == 0 {
            map_1[receiverIp] = UserInfo{Username: username, Chat: []string{msg}}
        } else {
            mydata := map_1[receiverIp]
            mydata.Username = username
            mydata.Chat = append(mydata.Chat, msg)
            map_1[receiverIp] = mydata
        }

        t, err := json.Marshal(map_1)
        if err != nil {
            panic(err)
        }

        finalOut := string(t)

        myfile, err := os.OpenFile("./mylogfile.txt", os.O_CREATE|os.O_WRONLY, 0644)
        if err != nil {
            fmt.Println(err)
            return
        }

        _, err = io.WriteString(myfile, finalOut)
        if err != nil {
            fmt.Println(err)
            return
        }

        myfile.Close()

        if err != nil {
            panic(err)
        }

        c.Close()

        break

    }
}

func SendMessage(message, host string) {

    CONNECT := Ip
    c, err := net.Dial("tcp", CONNECT)
    if err != nil {
        fmt.Println(err)
        return
    }

    c.Write([]byte("Shubham :" + " " + message + "\n"))
    bufio.NewReader(c).ReadString('\n')

}

