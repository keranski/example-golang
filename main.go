package main

import (
	"fmt"
	"log"
	"math/big"
	"net/http"

	"github.com/btcsuite/btcec"
	"github.com/btcsuite/btcnet"
	"github.com/btcsuite/btcutil"
)

func main() {
	port := os.Getenv("PORT")

	if port == "" {
		port = "8080"
	}

	http.HandleFunc("/", HelloHandler)

	log.Println("Listening on port", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), nil))
}

func HelloHandler(w http.ResponseWriter, _ *http.Request) {
	fmt.Fprintf(w, "Hello from Koyeb\n")
}
