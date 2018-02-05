package main

import ("fmt"
"sync")

var wg sync.WaitGroup

func timesFive(c chan int,n int){
  defer wg.Done()
  c <- n*5
}



func main(){
  num_channel := make(chan int)

  for i:=0 ; i < 10; i++ {
    wg.Add(1)  
    go timesFive(num_channel,i)

  }

  go func() {
    wg.Wait()
    close(num_channel)
  }()



  for item := range(num_channel){
    fmt.Println(item)
  }

}

