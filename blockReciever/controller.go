package main

import (
	"fmt"
	"net/http"
)

func homePage(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, "OK!")
	fmt.Println("Endpoint Hit: homePage")
}

func startApi() {
	http.HandleFunc("/", homePage)
	//log.Fatal(http.ListenAndServe(":1488", nil))
	go http.ListenAndServe(":1488", nil)
}