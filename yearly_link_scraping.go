package main

import (
  "fmt"
  //"github.com/PuerkitoBio/goquery"
)

var root_url = "http://spiderbites.nytimes.com/free_2014/index.html"

func scrape_page(l string) string{
  return l
}


func main(){
  fmt.Println(scrape_page(root_url))

}
