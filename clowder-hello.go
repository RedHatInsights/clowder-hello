package main

import (
	"fmt"
	"net/http"
	"os"

	client "github.com/redhatinsights/app-common-go/pkg/api/v1"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	msg := fmt.Sprintf("Args: %v", os.Args)
	fmt.Fprintln(w, msg)
}

func main() {
	if len(os.Args) > 1 {
		fmt.Printf("Hi")
	} else {
		port := client.LoadedConfig.WebPort
		http.HandleFunc("/", helloHandler)
		http.HandleFunc("/healthz", helloHandler)

		address := fmt.Sprintf(":%d", port)

		fmt.Printf("Started, serving at %s\n", address)
		err := http.ListenAndServe(address, nil)
		if err != nil {
			panic("ListenAndServe: " + err.Error())
		}
	}
}
