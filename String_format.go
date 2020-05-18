package main

import (
  "fmt"
)

func main(){
  var template=[5]string{"my","name","is","onkar","kulkarni"}
  Frameit(template)
}

func Frameit(template [5]string){
   maxlength := func(template [5]string) int{
    var temp int = 0
    for i:=0;i<len(template);i++{

        if len(template[i])>temp{
          temp=len(template[i])
          }
    }
    return temp
  }
  hlen:=maxlength(template)+2
  vlen:=len(template)+2


  for i:=0;i<vlen;i++{
    if i==0 || i==(vlen-1){
      fmt.Println("")
      for j:=0;j<hlen;j++{
        fmt.Print("*")
      }
      continue
    }
    noofstars:=(hlen-len(template[i-1]))/2
    fmt.Println()
    for j:=0;j<noofstars;j++{
        fmt.Print("*")
    }
    fmt.Print(template[i-1])
    for j:=0;j<noofstars;j++{
        fmt.Print("*")
    }
  }

}
