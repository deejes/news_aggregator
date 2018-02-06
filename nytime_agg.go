package main


import (
  "fmt"
  //"net/http"
  //"io/ioutil"
  //"golang.org/x/net/html"
  "github.com/PuerkitoBio/goquery"
  "strconv"
)


var root_url = "http://spiderbites.nytimes.com/"

func main(){
  //response,_ := http.Get("http://spiderbites.nytimes.com/")
  //doc,_ := goquery.NewDocumentFromReader(io.Reader(response.Body)
  doc,_ := goquery.NewDocument("http://spiderbites.nytimes.com/")
//  resp, _ := http.Get("")
 // z := html.NewTokenizer(resp.Body)
 // for i:=0;i<1000;i++{
 //   fmt.Println(z.Next())
 //   token := (z.Token())
 //   fmt.Println(string(token))
 //   }
  sel := doc.Find("h3").First().Parent().Contents()
  fmt.Println(sel.Text())
  for i := range sel.Nodes {
    node := sel.Eq(i)
    node_text, _ := strconv.Atoi(node.Text())

    if node_text >= 2000 {
    url,_ := node.Attr("href")
    
    fmt.Println(root_url+url)
    }
      // use `single` as a selection of 1 nod
    }
  }


