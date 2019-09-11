package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	hname, _ := os.Hostname()
	http.HandleFunc("/orgs/xendit/comments", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello from %s", hname)
	})

	http.ListenAndServe(fmt.Sprintf(":%s", os.Getenv("PORT")), nil)
}
