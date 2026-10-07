package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
	"sync"
)

// A tiny actor fixture: /readyz never changes state; POST / increments both
// memory and a durable file, so a cold restart cannot masquerade as a restore.
func main() {
	var mu sync.Mutex
	memory := 0
	http.HandleFunc("/readyz", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/" || r.Method != http.MethodPost {
			http.Error(w, "use POST /", http.StatusMethodNotAllowed)
			return
		}
		mu.Lock()
		defer mu.Unlock()
		file := 0
		data, err := os.ReadFile("/data/count")
		if err == nil {
			file, err = strconv.Atoi(strings.TrimSpace(string(data)))
		}
		if err != nil && !os.IsNotExist(err) {
			http.Error(w, "cannot read counter", http.StatusInternalServerError)
			return
		}
		if err := os.WriteFile("/data/count", []byte(strconv.Itoa(file+1)), 0600); err != nil {
			http.Error(w, "cannot persist counter", http.StatusInternalServerError)
			return
		}
		memory++
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]int{"memory": memory, "file": file + 1})
	})
	log.Fatal(http.ListenAndServe(":80", nil))
}
