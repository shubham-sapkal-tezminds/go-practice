package main

import (
	"fmt"
)

    
var userNames = map[string]string{
    "Naveen": "192.168.1.42:6000",
    "kaushal": "192.168.1.19:9000",
 }

var receiverIp string


    func main() {
    // Set properties of the predefined Logger, including
    // the log entry prefix and a flag to disable printing
    // the time, source file, and line number.
    // log.SetPrefix("greetings: ")
    // log.SetFlags(0)

    // Request a greeting message.
    // message, err := greetings.Hello("")
    // If an error was returned, print it to the console and
    // exit the program.
    // if err != nil {
    //     log.Fatal(err)
    // }

    // If no error was returned, print the returned message
    // to the console.
    // fmt.Println(message)

    fmt.Print("\nEnter username: ")
    var host string
    fmt.Scanln(&host)
    fmt.Println("name :",host)


    for key := range userNames { 
        // fmt.Printf("key[%s] value[%s]\n", k, v)
        if host == key {
       var user = userNames[host]

       fmt.Println("ip :-",user)

       receiverIp = user
    
        } 
    }

    something()

}

func something() {
    fmt.Println("assigned",receiverIp)
}