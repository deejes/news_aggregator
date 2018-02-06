package main


import (
  "fmt"
  //"net/http"
  //"io/ioutil"
  //"golang.org/x/net/html"
  "github.com/PuerkitoBio/goquery"
  //"strconv"
)


var root_url = "http://spiderbites.nytimes.com/"
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

//var year_url = "http://spiderbites.nytimes.com/free_2000/index.html"

//
//func main(){
//  var yearly_page_links []string
//  doc,_ := goquery.NewDocument(year_url)
//  sel := doc.Find("#mainContent").Find("li")
//  for i := range sel.Nodes{
//    node := sel.Eq(i).Children().First()
//    url,_ := node.Attr("href")
//    yearly_page_links= append(yearly_page_links, root_url+url)
//  }
//  fmt.Println(yearly_page_links)
//}


//var stories_links_page = "http://spiderbites.nytimes.com//free_2000/articles_2000_12_00002.html"
//func main(){
//  var story_links []string
//  doc,_ := goquery.NewDocument(stories_links_page)
//  sel := doc.Find("#headlines").Find("li")
//  for i := range sel.Nodes{
//    node := sel.Eq(i).Children().First()
//    url,_ := node.Attr("href")
//    story_links= append(story_links, url)
//  }
//  fmt.Println(story_links)
//}
//

type Article struct{
  Title string
  Author string
  Date string
  Body string
}

var article_page = "http://www.nytimes.com/2000/12/22/arts/art-review-party-time-inside-and-out-playful-wit-reigns-at-skidmore-s-new-museum.html"
func main(){
  var a Article
  doc,_ := goquery.NewDocument(article_page)

  title := doc.Find("h1")
  date := doc.Find(".dateline")
  author := doc.Find(".byline")
  body := doc.Find(".story-body-text")

  a.Title = title.Text()
  a.Author= author.Text()[3:]
  a.Date = date.Text()
  a.Body = body.Text()

  fmt.Println(a)
  //for i := range sel.Nodes{
  //node := sel.Eq(i).Children().First()
  //url,_ := node.Attr("href")
  //story_links= append(story_links, url)
  //}
}
