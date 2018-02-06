package main


import (
  "fmt"
  //"net/http"
  //"io/ioutil"
  //"golang.org/x/net/html"
  "github.com/PuerkitoBio/goquery"
  //"strconv"
)


//var root_url = "http://spiderbites.nytimes.com/"
//func main(){
//  var yearly_links []string
//  doc,_ := goquery.NewDocument(root_url)
//  sel := doc.Find("h3").First().Parent().Contents()
//  for i := range sel.Nodes {
//    node := sel.Eq(i)
//    node_text, _ := strconv.Atoi(node.Text())
//    if node_text >= 2000 {
//      url,_ := node.Attr("href")
//      yearly_links = append(yearly_links,root_url+url)
//    }
//  }
//  for ind,link := range yearly_links{
//  fmt.Println(ind,link)
//  }
//}

var root_url = "http://spiderbites.nytimes.com/free_2000/index.html"


func main(){
  var parts_links []string
  doc,_ := goquery.NewDocument(root_url)
  sel := doc.Find("#mainContent").Find("li")
  for i := range sel.Nodes{
    node := sel.Eq(i).Children().First()
    url,_ := node.Attr("href")
    fmt.Println(url)
    parts_links = append(parts_links,url)
  }
  fmt.Println(len(parts_links))
  fmt.Println(parts_links)
}
