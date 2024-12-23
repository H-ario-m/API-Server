package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"
)

type DataStore struct {
	data []string
}

func (ds *DataStore) Add(data string) {

	ds.data = append(ds.data, data)
}


func (ds *DataStore) GetAll() string {
	return strings.Join(ds.data, "\n")
}  

func main() {
	store := &DataStore{}

	http.HandleFunc("/store", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost {
			body, err := io.ReadAll(r.Body)
			if err != nil {
				http.Error(w, "Error reading request body", http.StatusInternalServerError)
				return
			}
			data := string(body)
			store.Add(data)

			w.WriteHeader(http.StatusOK)
			fmt.Fprintln(w, data)
		}else if r.Method == http.MethodGet {
    w.WriteHeader(http.StatusOK)
    fmt.Fprintln(w, store.GetAll())
}else {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})

	log.Println("Server is running on port 8080...")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
