package main

import (
  "fmt"
)

func main(){
  var start1 int = 0
  var end int = 5
  //fmt.Scanln(&end)
  var answer =getFib(start1,end-1)
  fmt.Println(answer)
}


func getFib(start1 int, end int )int{
                                                  //1,1,2,3,5,8,13,21
    var temp int=1
    var reduce int
    for i:=0;i<end;i++{
      reduce=start1+temp
      start1=temp
      temp=reduce
    }
    return temp
}
