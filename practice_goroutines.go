package main

import (
  "fmt"
)


func cleanup(){
  r := recover() 
  if r!= nil{
    fmt.Println("recoevered in cleanup",r)
  }
}

func hello(s string){

  defer cleanup()
  fmt.Println("hello" + s)
  for i:=0;i < 4; i++ {
    fmt.Println(i)

    if i == 2{
      panic ("oh no!")
    }
  }
}

func main() {
  hello("ss")
}
