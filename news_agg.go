package main

import ("fmt"
        "net/http"
        "io/ioutil"
      
      )


func main(){
    resp,_ := http.Get("https://www.washingtonpost.com/news-sitemap-index.xml")
    bytes,_ := ioutil.ReadAll(resp.Body)
    body_string := string(bytes)
    fmt.Println(body_string)
    resp.Body.Close()
}
