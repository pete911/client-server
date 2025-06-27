package main

import (
	"fmt"
	"net/http"
	"os"
)

func server(port int) {
	m := http.NewServeMux()
	m.HandleFunc("/message", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(`{"status": "health"}`))
	})

	addr := fmt.Sprintf("localhost:%d", port)
	fmt.Printf("server listening on %s\n", addr)
	fmt.Println(http.ListenAndServe(addr, m))
	os.Exit(1)
}
