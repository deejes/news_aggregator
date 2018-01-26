package main

import ("fmt"
        "net/http"
        "io/ioutil"
        "encoding/xml"
      )

type SitemapIndex struct {
    Locations []string `xml:"sitemap>loc"`
}

type News struct {
    Titles []string `xml:"url>news>title"`
    Keywords []string `xml:"url>news>keywords"`
    Locations []string `xml:"url>loc"`
  }

type NewsMap struct{
  Keyword string
  Location string
}


func main(){
    var s SitemapIndex
    var n News
    news_map := make(map[string]NewsMap) // this is map(dict) with key of type string and value of NewsMap
    resp,_ := http.Get("https://www.washingtonpost.com/news-sitemap-index.xml")
    bytes,_ := ioutil.ReadAll(resp.Body)
    resp.Body.Close()
    xml.Unmarshal(bytes,&s) // so now xml has been unpacked into s of type SitemapIndex

    for _,Location := range s.Locations{ // iterating over s
        resp,_ := http.Get(Location)  // we're doing a get on each link that we got from the sitemap
        bytes,_ := ioutil.ReadAll(resp.Body)
        xml.Unmarshal(bytes,&n) // now from each link, results have been unpacked into news struct.

        for idx,_ := range n.Titles{ // idx is simply the indexes of n.Titles
        news_map[n.Titles[idx]] = NewsMap{n.Keywords[idx],n.Locations[idx]} // creates a hashmap(?) with titles as keys and keyword and location as values
        }
    }

    for idx,data := range news_map{
    fmt.Println("\n",idx)
    fmt.Println("\n",data.Keyword)
    fmt.Println("\n",data.Location)
    }

}
