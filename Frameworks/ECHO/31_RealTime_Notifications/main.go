package main

import (
	"fmt"
	"log"
	"net/http"
	"sync"
)

var adminClients = make(map[chan string]bool) // Global channel for list of admin SE clients

var mu sync.Mutex

func adminSSE(w http.ResponseWriter, r *http.Request) {
	// 1. Set the headers
	w.Header().Set("Content-Type", "text/event-stream")
	w.Header().Set("Cache-Control", "no-cache")
	w.Header().Set("Connection", "keep-alive")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	flusher, ok := w.(http.Flusher)
	if !ok {
		http.Error(w, "Streaming unsupported", http.StatusInternalServerError)
		return
	}

	clientChan := make(chan string)
	mu.Lock()
	adminClients[clientChan] = true
	mu.Unlock()

	// Remove client when function ends
	defer func() {
		mu.Lock()
		delete(adminClients, clientChan)
		mu.Unlock()
		close(clientChan)
	}()

	for {
		msg, ok := <-clientChan
		if !ok {
			return
		}
		fmt.Fprintf(w, "Status %s \n\n", msg)
		flusher.Flush()
	}
}

func userLogin(w http.ResponseWriter, r *http.Request) {
	user := r.URL.Query().Get("user")
	if user == "" {
		http.Error(w, "User required", http.StatusBadRequest)
		return
	}
	mu.Lock()
	for clientChan := range adminClients {
		clientChan <- fmt.Sprintf("User %s Loggedin", user)
	}
	mu.Unlock()
	w.Write([]byte("Login Success"))
}

func main() {
	http.HandleFunc("/admin/events", adminSSE)
	http.HandleFunc("/user/login", userLogin)
	fmt.Println("Server is running...")
	log.Fatal(http.ListenAndServe(":8090", nil))
}
