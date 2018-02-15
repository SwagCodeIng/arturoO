package main

import (
  "fmt"
  "net/http"
  "log"
)

func main(){
  http.HandleFunc("/hello", handleHello)
  fmt.Println("serving on http://localhost:7777/hello")
  log.Fatal(http.ListenAndServe("localhost:7777", nil))
}

func handleHello(w http.ResponseWriter, req *http.Request){
  log.Println("Serving", req.URL)
  fmt.Fprintln(w, "Primer microservicio en GO")
}
