package main

import (
  "fmt"
)

func main(){
  var n int
  fmt.Scanln(&n)
  fmt.Println(fact(n))
}

func fact( n int) int{
  if n==0 || n==1{
    return 1
  }else{
    return n * fact(n-1)
  }
}
