package main

import "fmt"

func timesFive(c chan int,n int){
  c <- n*5
}


func main(){
  num_channel := make(chan int)
  go timesFive(num_channel, 0)
  go timesFive(num_channel, 5)
  go timesFive(num_channel, 10)
  x := <-num_channel
  y := <-num_channel
  z := <-num_channel
  fmt.Println(x,y,z)
  // HOW IS THE ORDERING HAPPENING THOUGH?
}

