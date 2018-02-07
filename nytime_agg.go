package main


import (
  "fmt"
  "net/http"
  //"io/ioutil"
  //"golang.org/x/net/html"
  "github.com/PuerkitoBio/goquery"
  //"strconv"
  "html/template"
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

type Article struct{
  Title string
  //Author string
  Date string
  Body string
}

var article_array []Article

func get_article_data(url string) Article{
  var a Article
  doc,_ := goquery.NewDocument(url)

  title := doc.Find("h1")
  date := doc.Find(".dateline")
  // TODO catch/try for authors
  //author := doc.Find(".byline")
  body := doc.Find(".story-body-text")

  a.Title = title.Text()
  //a.Author= author.Text()[3:]
  a.Date = date.Text()
  a.Body = body.Text()
  return a
  //fmt.Println(a)
}


var stories_links_page = "http://spiderbites.nytimes.com//free_2000/articles_2000_12_00002.html"

//func main(){
//  get_from_subparts_page(stories_links_page)
//}

// TODO implement articles page struct and display articles through html
func Get_from_subparts_page(w http.ResponseWriter, r *http.Request){
  var count int
  var articles []Article
  var story_links []string
  doc,_ := goquery.NewDocument(stories_links_page)
  sel := doc.Find("#headlines").Find("li")
  // append article links to story_links slice
  for i := range sel.Nodes{
    node := sel.Eq(i).Children().First()
    url,_ := node.Attr("href")
    story_links= append(story_links, url)
  }
  // iterate over article links
  for i := range story_links{
    l := story_links[i]
    articles = append(articles,get_article_data(l))
    count++
    if count > 3 {
      break
    }
  }
  t, _ := template.ParseFiles("nyt_html.html")
  t.Execute(w, articles)
  fmt.Println(articles)
}


