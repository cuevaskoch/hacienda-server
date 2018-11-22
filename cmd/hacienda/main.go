package main

import (
	"fmt"
	"log"
	"net/http"

	services "./services"
)

func main() {
	http.HandleFunc("/", handleRoot)
	http.HandleFunc("/ad-blocking", services.HandleAdBlockingStatus)
	log.Fatal(http.ListenAndServe(":31337", nil))
}

func handleRoot(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Welcome to Hacienda, where all things are possible.")
}
