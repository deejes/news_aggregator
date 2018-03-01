package main

import (
  "net/http"
  "fmt"
)


func indexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>This is Devansh's news aggregtor</h1>")
 }

func main() {
  http.HandleFunc("/", indexHandler)
  http.HandleFunc("/wp/", newsAggHandler)
  //http.HandleFunc("/nyt/", Get_from_subparts_page)
  http.ListenAndServe(":8080", nil) 
}
