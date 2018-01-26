package main

import (
    "fmt"
    "net/http"
    "html/template"
)

type NewsAggPage struct{
    Title string
    News string
}


func indexHandler(w http.ResponseWriter, r *http.Request){
    fmt.Fprintf(w, "<h1>HW</h1>")
}

func newsAggHandler(w http.ResponseWriter, r *http.Request){
    p := NewsAggPage{Title:"Title",News:"News"}
    t, _ := template.ParseFiles("htmltemplate.html")
    t.Execute(w,p)
  }

func main(){
      http.HandleFunc("/",indexHandler)
      http.HandleFunc("/agg/",newsAggHandler)
      http.ListenAndServe(":8000",nil)
    }
