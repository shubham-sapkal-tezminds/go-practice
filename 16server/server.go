package main

import (
	"fmt"
	"log"
	"net/http"
)


func formHandler(w http.ResponseWriter, r *http.Request) {

	err := r.ParseForm()

	if err != nil {
		fmt.Fprintf(w, "ParseForm() err: %v", err)
        return	
	}

	fmt.Fprintf(w, "Post request successful\n")

	name := r.FormValue("name")
	address := r.FormValue("address")

	fmt.Fprintf(w, "Name : %s\n", name)
	fmt.Fprintf(w, "Address : %s\n", address)

}


func helloHandler(w http.ResponseWriter, r *http.Request) {   //  ResponseWriter- To send response back, Request - it gives more info about request
	if r.URL.Path != "/hello" {
		fmt.Println(w, "404 not found", http.StatusNotFound)  // Gives error "404 not found" if path != /hello
	}

	if r.Method != "GET" {
		fmt.Println(w, "Method is not supported", http.StatusNotFound)  //Gives error if method != GET
	}

	fmt.Fprintf(w, "Hello!")  // sends "Hello!" message
}


func main() {

	fmt.Println("Building basic web server using net/http package")

	fmt.Printf("Starting server at port: 8080 \n")


/* We need to pass handlers to server, so server know how to respond to request and which request to accept */

    fileServer := http.FileServer(http.Dir("./static")) // New code
    

// HandleFunc adds route handlers to the server , takes 2 arguments, 1st - path to listen for request , 2nd - func/handler to respond correctly to request

    http.Handle("/", fileServer)
    http.HandleFunc("/form", formHandler)
    http.HandleFunc("/hello", helloHandler)



	err := http.ListenAndServe(":8080", nil)  // ListenAndServe method starts the server and listens to the specified server i.e. 8080

	checkErrNil(err)
}


/* checkErrNil :- checks if error is there ang logs it to "log" */

func checkErrNil(err error) {     
	if err != nil {
		log.Fatal(err)
	}
}