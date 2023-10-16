package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

func fetchMetadata(w http.ResponseWriter, req *http.Request) {
	if file, _, fileError := req.FormFile(`uploadfile`); fileError != nil {
		fmt.Println("Form file error")
	} else {
		if newFile, errNewFile := os.Create("C:\\golang\\target.txt"); errNewFile != nil {
			fmt.Println("Create file error")

		} else {

			defer file.Close()
			defer newFile.Close()

			if _, errCopy := io.Copy(newFile, file); errCopy != nil {
				fmt.Println("Copy error")
			}
		}
		json.NewEncoder(w).Encode("Success")
	}
	json.NewEncoder(w).Encode("Failed")

}

func main() {
	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/upload", fetchMetadata)
	//myRouter.HandleFunc("/senddata", fetchMetadata)
	log.Fatal(http.ListenAndServe(":10002", myRouter))
}
