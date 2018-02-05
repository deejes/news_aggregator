package main

import (
  "fmt"
  "time"
  "sync"
)

var wg sync.WaitGroup

func say(s string){
  defer wg.Done()
  for i:=0; i<3;i++{
    fmt.Println(s)
    time.Sleep(time.Millisecond*500)
      }
    }


    func main() {
      go say("Hey")
      wg.Add(1)
      go say("There")
      wg.Add(1)
      wg.Wait()
      go say("Hi")
      wg.Add(1)
      wg.Wait()
      fmt.Println("s")
    }

