package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gobuffalo/packr"
	"github.com/jrkoch/hacienda-server/internal/services"
)

func main() {
	http.HandleFunc("/", handleRoot)
	http.HandleFunc("/ad-blocking", services.HandleAdBlocking)
	log.Fatal(http.ListenAndServe(":31337", nil))
}

func handleRoot(w http.ResponseWriter, r *http.Request) {
	box := packr.NewBox("./html")
	html, err := box.FindString("index.html")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Fprint(w, html)
}
