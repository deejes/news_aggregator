package main

import (
  "fmt"
  "net/http"
  "html/template"
  "encoding/xml"
  "io/ioutil"
  "sync"
)

var wg sync.WaitGroup

type NewsMap struct {
  Keyword string
  Location string
}

type NewsAggPage struct {
  Title string
  News map[string]NewsMap
}

type Sitemapindex struct {
  Locations []string `xml:"sitemap>loc"`
}

type News struct {
  Titles []string `xml:"url>news>title"`
  Keywords []string `xml:"url>news>keywords"`
  Locations []string `xml:"url>loc"`
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
  fmt.Fprintf(w, "<h1>This is Devansh's news aggregtor</h1>")
}

func get_url_data(c chan News, url string){
  defer wg.Done()
  var n News
  resp, _ := http.Get(url)
  bytes, _ := ioutil.ReadAll(resp.Body)
  xml.Unmarshal(bytes, &n)
  resp.Body.Close()
  c <- n
}



func newsAggHandler(w http.ResponseWriter, r *http.Request) {
  var s Sitemapindex
  resp, _ := http.Get("https://www.washingtonpost.com/news-sitemap-index.xml")
  bytes, _ := ioutil.ReadAll(resp.Body)
  xml.Unmarshal(bytes, &s)
  news_map := make(map[string]NewsMap)
  resp.Body.Close()

  queue := make(chan News,30)

  for _, Location := range s.Locations {
    wg.Add(1)
    go get_url_data(queue, Location)
  }

  wg.Wait()
  close(queue)

  for elem := range queue {
  for idx, _ := range elem.Keywords {
      news_map[elem.Titles[idx]] = NewsMap{elem.Keywords[idx], elem.Locations[idx]}
    }
  }

  p := NewsAggPage{Title: "Current News from WP", News: news_map}
  t, _ := template.ParseFiles("htmltemplate.html")
  t.Execute(w, p)
}

func main() {
  http.HandleFunc("/", indexHandler)
  http.HandleFunc("/wp/", newsAggHandler)
  http.ListenAndServe(":8080", nil) 
}
