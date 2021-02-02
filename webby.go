package main 

import (
  "fmt"
  "log"
  "net/http"
)

func rootHandler(w http.ResponseWriter, r *http.Request) {
  path := r.URL.Path[1:] 
  fmt.Fprintf(w, "Path: %s", path)
  log.Printf("Got path: %s", path) 
}

func main() {
  log.Println("Serving on 8080")
  http.HandleFunc("/", rootHandler)
  log.Fatal(http.ListenAndServe(":8080", nil))
}
