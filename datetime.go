package main

import (
    "fmt"
    "time"
    //"strconv"
    )
func main() {

  currentdate := time.Now().Unix()

  weekago := currentdate - (7*24*60*60)
  //result :=strconv.FormatInt(weekago,10)
  result := time.Unix(weekago, 0)
  fmt.Println(result)

}
